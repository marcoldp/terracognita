package provider

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/chr4/pwgen"
	"github.com/cycloidio/terraforming/errcode"
	"github.com/cycloidio/terraforming/filter"
	"github.com/cycloidio/terraforming/tag"
	"github.com/cycloidio/terraforming/util"
	"github.com/cycloidio/terraforming/writer"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
)

// Resource represents the minimal information needed to
// define a Provider resource
type Resource struct {
	// ID is the ID of the Resource
	ID string

	// Type is the type of resource (ex: aws_instance)
	Type string

	// TFResource is the definition of that resource
	TFResource *schema.Resource

	// Data is the actual data of the Resource
	Data *schema.ResourceData

	// Provider is the Provider of that Resource
	Provider Provider
}

// Read read the remote information of the Resource
func (r *Resource) Read(f filter.Filter) error {
	// Retry if any error happen
	err := util.RetryDefault(func() error {
		return r.TFResource.Read(r.Data, r.Provider.TFClient())
	})
	if err != nil {
		return err
	}

	// For some reason it failed to fetch the Resource, it should not be an error
	// because it could be an account related resource that it's not delcared or
	// is default.
	// Therefore we just continue
	if r.Data.Id() == "" {
		return errors.Wrapf(errcode.ErrResourceNotRead, "the resource %q with ID %q did not return an ID", r.Type, r.ID)
	}

	// Some resources can not be filtered by tags,
	// so we have to do it manually
	// it's not all of them though
	for _, t := range f.Tags {
		if v, ok := r.Data.GetOk(fmt.Sprintf("tags.%s", t.Name)); ok && v.(string) != t.Value {
			return errors.WithStack(errcode.ErrResourceDoNotMatchTag)
		}
	}

	return nil
}

// State calculates the state of the Resource and
// writes it to w
func (r *Resource) State(w writer.Writer) error {
	if importer := r.TFResource.Importer; importer != nil {
		resourceDatas, err := importer.State(r.Data, r.Provider.TFClient())
		if err != nil {
			return err
		}
		// TODO: The multple return could potentially be the `depends_on` of the
		// terraform.ResourceState
		// Investigate on SG
		for i, rd := range resourceDatas {
			if i != 0 {
				// for now we scape all the other ones
				// the firs one is the one we need and
				// for what've see the others are
				// 'aws_security_group_rules' (in aws)
				// and are not imported by default by
				// Terraform
				continue
			}

			tis := rd.State()
			if tis == nil {
				// IDK why some times it does not have
				// the ID (the only case in tis it's nil)
				// so if nil we don't need it
				continue
			}
			trs := &terraform.ResourceState{
				Type:     r.Type,
				Primary:  tis,
				Provider: "aws",
			}

			err := w.Write(fmt.Sprintf("%s.%s", tis.Ephemeral.Type, tag.GetNameFromTag(r.Provider.TagKey(), rd, r.ID)), trs)
			if err != nil {
				if errors.Cause(err) == writer.ErrAlreadyExistsKey {
					err = w.Write(fmt.Sprintf("%s.%s", tis.Ephemeral.Type, pwgen.Alpha(5)), trs)
					if err != nil {
						return err
					}
					return nil
				}
				return err
			}
		}
	}
	return nil
}

// HCL returns the HCL configuration of the Resource and
// writes it to HCL
func (r *Resource) HCL(w writer.Writer) error {
	err := w.Write(fmt.Sprintf("%s.%s", r.Type, tag.GetNameFromTag(r.Provider.TagKey(), r.Data, r.ID)), mergeFullConfig(r.Data, r.TFResource.Schema, ""))
	if err != nil {
		if errors.Cause(err) == writer.ErrAlreadyExistsKey {
			err = w.Write(fmt.Sprintf("%s.%s", r.Type, pwgen.Alpha(5)), mergeFullConfig(r.Data, r.TFResource.Schema, ""))
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

// mergeFullConfig creates the key to the map and if it had a value before set it, if
func mergeFullConfig(cfgr *schema.ResourceData, sch map[string]*schema.Schema, key string) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range sch {
		// If it's just a Computed value, do not add it to the output
		if v.Computed && !v.Optional && !v.Required {
			continue
		}

		// Basically calculates the needed
		// key to the current access
		if key != "" {
			key = key + "." + k
		} else {
			key = k
		}

		// schema.Resource means that it has nested fields
		if sr, ok := v.Elem.(*schema.Resource); ok {
			// Example would be aws_security_group
			if v.Type == schema.TypeSet {
				s, ok := cfgr.GetOk(key)
				if !ok {
					continue
				}

				res[k] = normalizeSetList(sr.Schema, s.(*schema.Set).List())
			} else if v.Type == schema.TypeList {
				//var ar interface{}
				var ar interface{} = make([]interface{}, 0)

				l, ok := cfgr.GetOk(key)
				if !ok {
					continue
				}

				list := l.([]interface{})
				for i := range list {
					ar = append(ar.([]interface{}), mergeFullConfig(cfgr, sr.Schema, fmt.Sprintf("%s.%d", key, i)))
				}

				res[k] = ar
			} else {
				res[k] = mergeFullConfig(cfgr, sr.Schema, key)
			}
			// As it's a nested element it does not require any of
			// the other code as it's for singel value schemas
			continue
		}

		vv, ok := cfgr.GetOk(key)
		// If the value is Required we need to add it
		// even if it's not send
		if (!ok || vv == nil) && !v.Required {
			continue
		}

		if s, ok := vv.(*schema.Set); ok {
			res[k] = s.List()
		} else {
			res[k] = normalizeInterpolation(normalizeValue(vv))
		}
	}
	return res
}

// normalizeValue removes the \n from the value now
func normalizeValue(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return strings.Replace(s, "\n", "", -1)
	}
	return v
}

var iamInternpolationRe = regexp.MustCompile(`(\$\{[^}]+\})`)

// normalizeInterpolation fixes the https://github.com/hashicorp/terraform/issues/18937
// on reading
func normalizeInterpolation(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return iamInternpolationRe.ReplaceAllString(s, `$$$1`)
	}
	return v
}

// normalizeSetList returns the normalization of a schema.Set.List
// it could be a simple list or a embedded structure.
// The sch it's used to also add required values if needed
func normalizeSetList(sch map[string]*schema.Schema, list []interface{}) interface{} {
	var ar interface{} = make([]interface{}, 0)

	for _, set := range list {
		switch val := set.(type) {
		case map[string]interface{}:
			// This case it's when a TypeSet has
			// a nested structure,
			// example: aws_security_group.ingress
			res := make(map[string]interface{})
			for k, v := range val {
				switch vv := v.(type) {
				case *schema.Set:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv.List())
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case []interface{}:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv)
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case interface{}:
					if !isDefault(sch[k], v) {
						res[k] = v
					}
				}
			}
			ar = append(ar.([]interface{}), res)
		case []interface{}:
			ns := normalizeSetList(sch, val)
			if !isDefault(nil, ns) {
				ar = append(ar.([]interface{}), ns)
			}
		case interface{}:
			// This case is normally for the
			// "Type: schema.TypeSet, Elm: schema.Schema{Type: schema.TypeString}"
			// definitions on TF,
			// example: aws_security_group.ingress.security_groups
			if !isDefault(nil, val) {
				ar = append(ar.([]interface{}), val)
			}
		}
	}

	return ar
}

var (
	// Ideally this could be generated using "enumer", it
	// would be a better idea as then we do not have
	// to maintain this list
	tfTypes = []schema.ValueType{
		schema.TypeBool,
		schema.TypeInt,
		schema.TypeFloat,
		schema.TypeString,
		schema.TypeList,
		schema.TypeMap,
		schema.TypeSet,
	}
)

// isDefault is used on normalizSet as the Sets do not use the normal
// TF strucure (access by key) and are stored as raw maps with some
// default values that we don't want on the HCL output.
// example: [], false, "", 0 ...
func isDefault(sch *schema.Schema, v interface{}) bool {
	if sch != nil && sch.Required {
		return false
	}

	for _, t := range tfTypes {
		if reflect.DeepEqual(t.Zero(), v) {
			return true
		}
	}
	return false
}
