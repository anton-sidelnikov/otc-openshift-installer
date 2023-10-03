package defaults

import (
	"testing"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	openstackdefaults "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack/defaults"
	"github.com/stretchr/testify/assert"
)

func defaultInstallConfig() *types.InstallConfig {
	return &types.InstallConfig{
		AdditionalTrustBundlePolicy: defaultAdditionalTrustBundlePolicy(),
		Networking: &types.Networking{
			MachineNetwork: []types.MachineNetworkEntry{
				{CIDR: *DefaultMachineCIDR},
			},
			NetworkType:    defaultNetworkType,
			ServiceNetwork: []ipnet.IPNet{*defaultServiceNetwork},
			ClusterNetwork: []types.ClusterNetworkEntry{
				{
					CIDR:       *defaultClusterNetwork,
					HostPrefix: int32(defaultHostPrefix),
				},
			},
		},
		ControlPlane: defaultMachinePool("master"),
		Compute:      []types.MachinePool{*defaultMachinePool("worker")},
		Publish:      types.ExternalPublishingStrategy,
	}
}

func defaultInstallConfigWithEdge() *types.InstallConfig {
	c := defaultInstallConfig()
	c.Compute = append(c.Compute, *defaultMachinePool("edge"))
	return c
}

func defaultOpenStackInstallConfig() *types.InstallConfig {
	c := defaultInstallConfig()
	c.Platform.OpenStack = &openstack.Platform{}
	openstackdefaults.SetPlatformDefaults(c.Platform.OpenStack, c.Networking)
	return c
}

func defaultAdditionalTrustBundlePolicy() types.PolicyType {
	return types.PolicyProxyOnly
}

func TestSetInstallConfigDefaults(t *testing.T) {
	cases := []struct {
		name     string
		config   *types.InstallConfig
		expected *types.InstallConfig
	}{
		{
			name:     "empty",
			config:   &types.InstallConfig{},
			expected: defaultInstallConfig(),
		},
		{
			name: "empty OpenStack",
			config: &types.InstallConfig{
				Platform: types.Platform{
					OpenStack: &openstack.Platform{},
				},
			},
			expected: defaultOpenStackInstallConfig(),
		},
		{
			name: "Networking present",
			config: &types.InstallConfig{
				Networking: &types.Networking{},
			},
			expected: defaultInstallConfig(),
		},
		{
			name: "Networking types present",
			config: &types.InstallConfig{
				Networking: &types.Networking{
					NetworkType: "test-networking-type",
				},
			},
			expected: func() *types.InstallConfig {
				c := defaultInstallConfig()
				c.Networking.NetworkType = "test-networking-type"
				return c
			}(),
		},
		{
			name: "Service network present",
			config: &types.InstallConfig{
				Networking: &types.Networking{
					ServiceNetwork: []ipnet.IPNet{*ipnet.MustParseCIDR("1.2.3.4/8")},
				},
			},
			expected: func() *types.InstallConfig {
				c := defaultInstallConfig()
				c.Networking.ServiceNetwork[0] = *ipnet.MustParseCIDR("1.2.3.4/8")
				return c
			}(),
		},
		{
			name: "Cluster network present",
			config: &types.InstallConfig{
				Networking: &types.Networking{
					ClusterNetwork: []types.ClusterNetworkEntry{
						{
							CIDR:       *ipnet.MustParseCIDR("8.8.0.0/18"),
							HostPrefix: 22,
						},
					},
				},
			},
			expected: func() *types.InstallConfig {
				c := defaultInstallConfig()
				c.Networking.ClusterNetwork = []types.ClusterNetworkEntry{
					{
						CIDR:       *ipnet.MustParseCIDR("8.8.0.0/18"),
						HostPrefix: 22,
					},
				}
				return c
			}(),
		},
		{
			name: "control plane present",
			config: &types.InstallConfig{
				ControlPlane: &types.MachinePool{},
			},
			expected: defaultInstallConfig(),
		},
		{
			name: "Compute present",
			config: &types.InstallConfig{
				Compute: []types.MachinePool{{Name: "worker"}},
			},
			expected: func() *types.InstallConfig {
				c := defaultInstallConfig()
				c.Compute = []types.MachinePool{*defaultMachinePool("worker")}
				return c
			}(),
		},
		{
			name: "Edge Compute present",
			config: &types.InstallConfig{
				Compute: []types.MachinePool{{Name: "worker"}, {Name: "edge"}},
			},
			expected: func() *types.InstallConfig {
				c := defaultInstallConfigWithEdge()
				c.Compute = []types.MachinePool{
					*defaultMachinePool("worker"),
					*defaultEdgeMachinePool("edge"),
				}
				return c
			}(),
		},
		{
			name: "OpenStack platform present",
			config: &types.InstallConfig{
				Platform: types.Platform{
					OpenStack: &openstack.Platform{},
				},
			},
			expected: func() *types.InstallConfig {
				c := defaultOpenStackInstallConfig()
				return c
			}(),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SetInstallConfigDefaults(tc.config)
			assert.Equal(t, tc.expected, tc.config, "unexpected install config")
		})
	}
}
