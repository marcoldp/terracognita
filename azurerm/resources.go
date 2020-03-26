package azurerm

import (
	"context"

	"github.com/pkg/errors"

	"github.com/cycloidio/terracognita/provider"
	"github.com/cycloidio/terracognita/tag"
)

// ResourceType is the type used to define all the Resources
// from the Provider
type ResourceType int

//go:generate enumer -type ResourceType -addprefix azurerm_ -transform snake -linecomment
const (
	ResourceGroup ResourceType = iota
	VirtualMachine
	VirtualNetwork
	Subnet
	NetworkInterface
	NetworkSecurityGroup
)

type rtFn func(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error)

var (
	resources = map[ResourceType]rtFn{
		ResourceGroup:        resourceGroup,
		VirtualMachine:       virtualMachines,
		VirtualNetwork:       virtualNetworks,
		Subnet:               subnets,
		NetworkInterface:     networkInterfaces,
		NetworkSecurityGroup: networkSecurityGroups,
	}
)

func resourceGroup(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	resourceGroup := a.azurer.GetResourceGroup()
	r := provider.NewResource(*resourceGroup.ID, resourceType, a)
	resources := []provider.Resource{r}
	return resources, nil
}

func virtualMachines(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	virtualMachines, err := a.azurer.ListVirtualMachines(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list virtual machines from reader")
	}
	resources := make([]provider.Resource, 0, len(virtualMachines))
	for _, virtualMachine := range virtualMachines {
		r := provider.NewResource(*virtualMachine.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}

func virtualNetworks(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	virtualNetworks, err := a.azurer.ListVirtualNetworks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list virtual networks from reader")
	}
	resources := make([]provider.Resource, 0, len(virtualNetworks))
	for _, virtualNetwork := range virtualNetworks {
		r := provider.NewResource(*virtualNetwork.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}

func subnets(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	virtualNetworks, err := a.azurer.ListVirtualNetworks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list virtual networks from reader")
	}
	resources := make([]provider.Resource, 0)
	for _, virtualNetwork := range virtualNetworks {
		subnets, err := a.azurer.ListSubnets(ctx, *virtualNetwork.Name)
		if err != nil {
			return nil, errors.Wrap(err, "unable to list subnets from reader")
		}
		for _, subnet := range subnets {
			r := provider.NewResource(*subnet.ID, resourceType, a)
			resources = append(resources, r)
		}
	}
	return resources, nil
}

func networkInterfaces(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	networkInterfaces, err := a.azurer.ListInterfaces(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list network interfaces from reader")
	}
	resources := make([]provider.Resource, 0, len(networkInterfaces))
	for _, networkInterface := range networkInterfaces {
		r := provider.NewResource(*networkInterface.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}

func networkSecurityGroups(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	securityGroups, err := a.azurer.ListSecurityGroups(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list network security groups from reader")
	}
	resources := make([]provider.Resource, 0, len(securityGroups))
	for _, securityGroup := range securityGroups {
		r := provider.NewResource(*securityGroup.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}
