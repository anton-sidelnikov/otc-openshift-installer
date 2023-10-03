package defaults

import (
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	openstackdefaults "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack/defaults"
	operv1 "github.com/openshift/api/operator/v1"
)

var (
	// DefaultMachineCIDR default machine CIDR applied to MachineNetwork.
	DefaultMachineCIDR    = ipnet.MustParseCIDR("10.0.0.0/16")
	defaultServiceNetwork = ipnet.MustParseCIDR("172.30.0.0/16")
	defaultClusterNetwork = ipnet.MustParseCIDR("10.128.0.0/14")
	defaultHostPrefix     = 23
	defaultNetworkType    = string(operv1.NetworkTypeOVNKubernetes)
)

// SetInstallConfigDefaults sets the defaults for the install config.
func SetInstallConfigDefaults(c *types.InstallConfig) {
	if c.Networking == nil {
		c.Networking = &types.Networking{}
	}
	if len(c.Networking.MachineNetwork) == 0 {
		c.Networking.MachineNetwork = []types.MachineNetworkEntry{
			{CIDR: *DefaultMachineCIDR},
		}
	}
	if c.Networking.NetworkType == "" {
		c.Networking.NetworkType = defaultNetworkType
	}
	if len(c.Networking.ServiceNetwork) == 0 {
		c.Networking.ServiceNetwork = []ipnet.IPNet{*defaultServiceNetwork}
	}
	if len(c.Networking.ClusterNetwork) == 0 {
		c.Networking.ClusterNetwork = []types.ClusterNetworkEntry{
			{
				CIDR:       *defaultClusterNetwork,
				HostPrefix: int32(defaultHostPrefix),
			},
		}
	}

	if c.Publish == "" {
		c.Publish = types.ExternalPublishingStrategy
	}

	if c.ControlPlane == nil {
		c.ControlPlane = &types.MachinePool{}
	}
	c.ControlPlane.Name = "master"
	SetMachinePoolDefaults(c.ControlPlane, c.Platform.Name())

	defaultComputePoolUndefined := true
	for _, compute := range c.Compute {
		if compute.Name == types.MachinePoolComputeRoleName {
			defaultComputePoolUndefined = false
			break
		}
	}
	if defaultComputePoolUndefined {
		c.Compute = append(c.Compute, types.MachinePool{Name: types.MachinePoolComputeRoleName})
	}
	for i := range c.Compute {
		SetMachinePoolDefaults(&c.Compute[i], c.Platform.Name())
	}

	switch {
	case c.Platform.OpenStack != nil:
		openstackdefaults.SetPlatformDefaults(c.Platform.OpenStack, c.Networking)
	}

	if c.AdditionalTrustBundlePolicy == "" {
		c.AdditionalTrustBundlePolicy = types.PolicyProxyOnly
	}
}
