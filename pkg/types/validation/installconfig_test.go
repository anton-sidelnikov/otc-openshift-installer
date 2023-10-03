package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
	utilsslice "k8s.io/utils/strings/slices"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	configv1 "github.com/openshift/api/config/v1"
)

const TechPreviewNoUpgrade = "TechPreviewNoUpgrade"

func validInstallConfig() *types.InstallConfig {
	return &types.InstallConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: types.InstallConfigVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cluster",
		},
		BaseDomain:   "test-domain",
		Networking:   validIPv4NetworkingConfig(),
		ControlPlane: validMachinePool("master"),
		Compute:      []types.MachinePool{*validMachinePool("worker")},
		Platform: types.Platform{
			OpenStack: validOpenStackPlatform(),
		},
		PullSecret: `{"auths":{"example.com":{"auth":"authorization value"}}}`,
		Publish:    types.ExternalPublishingStrategy,
		Proxy: &types.Proxy{
			HTTPProxy:  "http://user:password@127.0.0.1:8080",
			HTTPSProxy: "https://user:password@127.0.0.1:8080",
			NoProxy:    "valid-proxy.com,172.30.0.0/16",
		},
	}
}

func validOpenStackPlatform() *openstack.Platform {
	return &openstack.Platform{
		Cloud:           "test-cloud",
		ExternalNetwork: "test-network",
		DefaultMachinePlatform: &openstack.MachinePool{
			FlavorName: "test-flavor",
		},
		APIVIPs:     []string{"10.0.0.5"},
		IngressVIPs: []string{"10.0.0.4"},
	}
}

func validIPv4NetworkingConfig() *types.Networking {
	return &types.Networking{
		NetworkType: "OVNKubernetes",
		MachineNetwork: []types.MachineNetworkEntry{
			{
				CIDR: *ipnet.MustParseCIDR("10.0.0.0/16"),
			},
		},
		ServiceNetwork: []ipnet.IPNet{
			*ipnet.MustParseCIDR("172.30.0.0/16"),
		},
		ClusterNetwork: []types.ClusterNetworkEntry{
			{
				CIDR:       *ipnet.MustParseCIDR("192.168.1.0/24"),
				HostPrefix: 28,
			},
		},
	}
}

func validIPv6NetworkingConfig() *types.Networking {
	return &types.Networking{
		NetworkType: "OVNKubernetes",
		MachineNetwork: []types.MachineNetworkEntry{
			{
				CIDR: *ipnet.MustParseCIDR("ffd0::/48"),
			},
		},
		ServiceNetwork: []ipnet.IPNet{
			*ipnet.MustParseCIDR("ffd1::/112"),
		},
		ClusterNetwork: []types.ClusterNetworkEntry{
			{
				CIDR:       *ipnet.MustParseCIDR("ffd2::/48"),
				HostPrefix: 64,
			},
		},
	}
}

func validDualStackNetworkingConfig() *types.Networking {
	return &types.Networking{
		NetworkType: "OVNKubernetes",
		MachineNetwork: []types.MachineNetworkEntry{
			{
				CIDR: *ipnet.MustParseCIDR("10.0.0.0/16"),
			},
			{
				CIDR: *ipnet.MustParseCIDR("ffd0::/48"),
			},
		},
		ServiceNetwork: []ipnet.IPNet{
			*ipnet.MustParseCIDR("172.30.0.0/16"),
			*ipnet.MustParseCIDR("ffd1::/112"),
		},
		ClusterNetwork: []types.ClusterNetworkEntry{
			{
				CIDR:       *ipnet.MustParseCIDR("192.168.1.0/24"),
				HostPrefix: 28,
			},
			{
				CIDR:       *ipnet.MustParseCIDR("ffd2::/48"),
				HostPrefix: 64,
			},
		},
	}
}

func TestValidateInstallConfig(t *testing.T) {
	cases := []struct {
		name          string
		installConfig *types.InstallConfig
		expectedError string
	}{
		{
			name:          "minimal",
			installConfig: validInstallConfig(),
		},
		{
			name: "invalid version",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.APIVersion = "bad-version"
				return c
			}(),
			expectedError: fmt.Sprintf(`^apiVersion: Invalid value: "bad-version": install-config version must be %q`, types.InstallConfigVersion),
		},
		{
			name: "invalid name",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ObjectMeta.Name = "bad-name-"
				return c
			}(),
			expectedError: `^metadata.name: Invalid value: "bad-name-": a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '\.', and must start and end with an alphanumeric character \(e\.g\. 'example\.com', regex used for validation is '\[a-z0-9]\(\[-a-z0-9]\*\[a-z0-9]\)\?\(\\\.\[a-z0-9]\(\[-a-z0-9]\*\[a-z0-9]\)\?\)\*'\)$`,
		},
		{
			name: "invalid ssh key",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.SSHKey = "bad-ssh-key"
				return c
			}(),
			expectedError: `^sshKey: Invalid value: "bad-ssh-key": ssh: no key found$`,
		},
		{
			name: "invalid base domain",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.BaseDomain = ".bad-domain."
				return c
			}(),
			expectedError: `^baseDomain: Invalid value: "\.bad-domain\.": a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '\.', and must start and end with an alphanumeric character \(e\.g\. 'example\.com', regex used for validation is '\[a-z0-9]\(\[-a-z0-9]\*\[a-z0-9]\)\?\(\\\.\[a-z0-9]\(\[-a-z0-9]\*\[a-z0-9]\)\?\)\*'\)$`,
		},
		{
			name: "overly long cluster domain",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ObjectMeta.Name = fmt.Sprintf("test-cluster%042d", 0)
				c.BaseDomain = fmt.Sprintf("test-domain%056d.a%060d.b%060d.c%060d", 0, 0, 0, 0)
				return c
			}(),
			expectedError: `^baseDomain: Invalid value: "` + fmt.Sprintf("test-cluster%042d.test-domain%056d.a%060d.b%060d.c%060d", 0, 0, 0, 0, 0) + `": must be no more than 253 characters$`,
		},
		{
			name: "missing networking",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking = nil
				return c
			}(),
			expectedError: `^networking: Required value: networking is required$`,
		},
		{
			name: "invalid network type",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.NetworkType = ""
				return c
			}(),
			expectedError: `^networking.networkType: Required value: network provider type required$`,
		},
		{
			name: "missing service network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ServiceNetwork = nil
				return c
			}(),
			expectedError: `^networking\.serviceNetwork: Required value: a service network is required$`,
		},
		{
			name: "invalid service network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ServiceNetwork[0] = *ipnet.MustParseCIDR("13.0.128.0/16")
				return c
			}(),
			expectedError: `^networking\.serviceNetwork\[0\]: Invalid value: "13\.0\.128\.0/16": invalid network address. got 13\.0\.128\.0/16, expecting 13\.0\.0\.0/16$`,
		},
		{
			name: "overlapping service network and machine cidr",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ServiceNetwork[0] = *ipnet.MustParseCIDR("10.0.2.0/24")
				return c
			}(),
			expectedError: `^networking\.serviceNetwork\[0\]: Invalid value: "10\.0\.2\.0/24": service network must not overlap with any of the machine networks$`,
		},
		{
			name: "overlapping machine network and machine network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.MachineNetwork = []types.MachineNetworkEntry{
					{CIDR: *ipnet.MustParseCIDR("13.0.0.0/16")},
					{CIDR: *ipnet.MustParseCIDR("13.0.2.0/24")},
				}

				return c
			}(),
			// also triggers the only-one-machine-network validation
			expectedError: `^networking\.machineNetwork\[1\]: Invalid value: "13\.0\.2\.0/24": machine network must not overlap with machine network 0$`,
		},
		{
			name: "overlapping service network and service network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ServiceNetwork = []ipnet.IPNet{
					*ipnet.MustParseCIDR("13.0.0.0/16"),
					*ipnet.MustParseCIDR("13.0.2.0/24"),
				}

				return c
			}(),
			// also triggers the only-one-service-network validation
			expectedError: `^\[networking\.serviceNetwork\[1\]: Invalid value: "13\.0\.2\.0/24": service network must not overlap with service network 0, networking\.serviceNetwork: Invalid value: "13\.0\.0\.0/16, 13\.0\.2\.0/24": only one service network can be specified]$`,
		},
		{
			name: "missing machine networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.MachineNetwork = nil
				return c
			}(),
			expectedError: `^networking\.machineNetwork: Required value: at least one machine network is required$`,
		},
		{
			name: "invalid machine cidr",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.MachineNetwork = []types.MachineNetworkEntry{{CIDR: *ipnet.MustParseCIDR("11.0.128.0/16")}}
				return c
			}(),
			expectedError: `^networking\.machineNetwork\[0\]: Invalid value: "11\.0\.128\.0/16": invalid network address. got 11\.0\.128\.0/16, expecting 11\.0\.0\.0/16$`,
		},
		{
			name: "invalid cluster network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ClusterNetwork = []types.ClusterNetworkEntry{{CIDR: *ipnet.MustParseCIDR("12.0.128.0/16"), HostPrefix: 23}}
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[0]\.cidr: Invalid value: "12\.0\.128\.0/16": invalid network address. got 12\.0\.128\.0/16, expecting 12\.0\.0\.0/16$`,
		},
		{
			name: "overlapping cluster network and machine cidr",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ClusterNetwork[0].CIDR = *ipnet.MustParseCIDR("10.0.3.0/24")
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[0]\.cidr: Invalid value: "10\.0\.3\.0/24": cluster network must not overlap with any of the machine networks$`,
		},
		{
			name: "overlapping cluster network and service network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ClusterNetwork[0].CIDR = *ipnet.MustParseCIDR("172.30.2.0/24")
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[0]\.cidr: Invalid value: "172\.30\.2\.0/24": cluster network must not overlap with service network 0$`,
		},
		{
			name: "overlapping cluster network and cluster network",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ClusterNetwork = []types.ClusterNetworkEntry{
					{CIDR: *ipnet.MustParseCIDR("12.0.0.0/16"), HostPrefix: 23},
					{CIDR: *ipnet.MustParseCIDR("12.0.3.0/24"), HostPrefix: 25},
				}
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[1]\.cidr: Invalid value: "12\.0\.3\.0/24": cluster network must not overlap with cluster network 0$`,
		},
		{
			name: "cluster network host prefix too large",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.ClusterNetwork[0].CIDR = *ipnet.MustParseCIDR("192.168.1.0/24")
				c.Networking.ClusterNetwork[0].HostPrefix = 23
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[0]\.hostPrefix: Invalid value: 23: cluster network host subnetwork prefix must not be larger size than CIDR 192.168.1.0/24$`,
		},
		{
			name: "cluster network host prefix unset",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.NetworkType = "OVNKubernetes"
				c.Networking.ClusterNetwork[0].CIDR = *ipnet.MustParseCIDR("192.168.1.0/24")
				c.Networking.ClusterNetwork[0].HostPrefix = 0
				return c
			}(),
			expectedError: `^networking\.clusterNetwork\[0]\.hostPrefix: Invalid value: 0: cluster network host subnetwork prefix must not be larger size than CIDR 192.168.1.0/24$`,
		},
		{
			name: "cluster network host prefix unset ignored",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.NetworkType = "HostPrefixNotRequiredPlugin"
				c.Networking.ClusterNetwork[0].CIDR = *ipnet.MustParseCIDR("192.168.1.0/24")
				return c
			}(),
			expectedError: ``,
		},
		{
			name: "missing control plane",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ControlPlane = nil
				return c
			}(),
			expectedError: `^controlPlane: Required value: controlPlane is required$`,
		},
		{
			name: "control plane with 0 replicas",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ControlPlane.Replicas = pointer.Int64Ptr(0)
				return c
			}(),
			expectedError: `^controlPlane.replicas: Invalid value: 0: number of control plane replicas must be positive$`,
		},
		{
			name: "invalid control plane",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ControlPlane.Replicas = nil
				return c
			}(),
			expectedError: `^controlPlane.replicas: Required value: replicas is required$`,
		},
		{
			name: "missing compute",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute = nil
				return c
			}(),
		},
		{
			name: "empty compute",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute = []types.MachinePool{}
				return c
			}(),
		},
		{
			name: "duplicate compute",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute = []types.MachinePool{
					*validMachinePool("worker"),
					*validMachinePool("worker"),
				}
				return c
			}(),
			expectedError: `^compute\[1\]\.name: Duplicate value: "worker"$`,
		},
		{
			name: "missing platform",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Platform = types.Platform{}
				return c
			}(),
			expectedError: `^platform: Invalid value: "": must specify one of the platforms \(alibabacloud, aws, azure, baremetal, external, gcp, ibmcloud, none, nutanix, openstack, powervs, vsphere\)$`,
		},
		{
			name: "valid openstack platform",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Platform = types.Platform{
					OpenStack: validOpenStackPlatform(),
				}
				return c
			}(),
		},
		{
			name: "empty proxy settings",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = ""
				c.Proxy.HTTPSProxy = ""
				c.Proxy.NoProxy = ""
				return c
			}(),
			expectedError: `^proxy: Required value: must include httpProxy or httpsProxy$`,
		},
		{
			name: "invalid HTTPProxy",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://bad%20uri"
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http://bad%20uri": parse "http://bad%20uri": invalid URL escape "%20"$`,
		},
		{
			name: "invalid HTTPProxy Schema missing",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http//baduri"
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http//baduri": parse "http//baduri": invalid URI for request$`,
		},
		{
			name: "HTTPProxy with port overlapping with Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://192.168.1.25:3030"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http://192.168.1.25:3030": proxy value is part of the cluster networks$`,
		},
		{
			name: "overlapping HTTPProxy and Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://192.168.1.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http://192.168.1.25": proxy value is part of the cluster networks$`,
		},
		{
			name: "non-overlapping HTTPProxy and Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://192.169.1.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
		},
		{
			name: "overlapping HTTPProxy and more than one Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://192.168.1.25"
				c.Networking = validIPv4NetworkingConfig()
				c.ClusterNetwork = append(c.ClusterNetwork, []types.ClusterNetworkEntry{
					{
						CIDR:       *ipnet.MustParseCIDR("192.168.0.0/16"),
						HostPrefix: 28,
					},
				}...,
				)
				return c
			}(),
			expectedError: `^\Q[networking.clusterNetwork[1].cidr: Invalid value: "192.168.0.0/16": cluster network must not overlap with cluster network 0, proxy.httpProxy: Invalid value: "http://192.168.1.25": proxy value is part of the cluster networks]\E$`,
		},
		{
			name: "non-overlapping HTTPProxy and Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://172.31.0.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
		},
		{
			name: "HTTPProxy with port overlapping with Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://172.30.0.25:3030"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http://172.30.0.25:3030": proxy value is part of the service networks$`,
		},
		{
			name: "overlapping HTTPProxy and Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://172.30.0.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpProxy: Invalid value: "http://172.30.0.25": proxy value is part of the service networks$`,
		},
		{
			name: "overlapping HTTPProxy and more than one Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "http://172.30.0.25"
				c.Networking = validIPv4NetworkingConfig()
				c.ServiceNetwork = append(c.ServiceNetwork, []ipnet.IPNet{
					*ipnet.MustParseCIDR("172.30.1.0/24"),
				}...,
				)
				return c
			}(),
			expectedError: `^\Q[networking.serviceNetwork[1]: Invalid value: "172.30.1.0/24": service network must not overlap with service network 0, networking.serviceNetwork: Invalid value: "172.30.0.0/16, 172.30.1.0/24": only one service network can be specified, proxy.httpProxy: Invalid value: "http://172.30.0.25": proxy value is part of the service networks]\E$`,
		},
		{
			name: "non-overlapping HTTPSProxy and Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://192.168.2.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
		},
		{
			name: "HTTPSProxy with port overlapping with Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://192.168.1.25:3030"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "http://192.168.1.25:3030": proxy value is part of the cluster networks$`,
		},
		{
			name: "overlapping HTTPSProxy and Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://192.168.1.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "http://192.168.1.25": proxy value is part of the cluster networks$`,
		},
		{
			name: "overlapping HTTPSProxy and more than one Cluster Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://192.168.1.25"
				c.Networking = validIPv4NetworkingConfig()
				c.ClusterNetwork = append(c.ClusterNetwork, []types.ClusterNetworkEntry{
					{
						CIDR:       *ipnet.MustParseCIDR("192.168.0.0/16"),
						HostPrefix: 28,
					},
				}...,
				)
				return c
			}(),
			expectedError: `^\Q[networking.clusterNetwork[1].cidr: Invalid value: "192.168.0.0/16": cluster network must not overlap with cluster network 0, proxy.httpsProxy: Invalid value: "http://192.168.1.25": proxy value is part of the cluster networks]\E$`,
		},
		{
			name: "overlapping HTTPSProxy and Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://172.30.0.25"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "http://172.30.0.25": proxy value is part of the service networks$`,
		},
		{
			name: "HTTPSProxy with port overlapping with Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://172.30.0.25:3030"
				c.Networking = validIPv4NetworkingConfig()
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "http://172.30.0.25:3030": proxy value is part of the service networks$`,
		},
		{
			name: "overlapping HTTPSProxy and more than one Service Networks",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http://172.30.0.25"
				c.Networking = validIPv4NetworkingConfig()
				c.ServiceNetwork = append(c.ServiceNetwork, []ipnet.IPNet{
					*ipnet.MustParseCIDR("172.30.1.0/24"),
				}...,
				)
				return c
			}(),
			expectedError: `^\Q[networking.serviceNetwork[1]: Invalid value: "172.30.1.0/24": service network must not overlap with service network 0, networking.serviceNetwork: Invalid value: "172.30.0.0/16, 172.30.1.0/24": only one service network can be specified, proxy.httpsProxy: Invalid value: "http://172.30.0.25": proxy value is part of the service networks]\E$`,
		},
		{
			name: "invalid HTTPProxy Schema different schema",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPProxy = "ftp://baduri"
				return c
			}(),
			expectedError: `^proxy.httpProxy: Unsupported value: "ftp": supported values: "http"$`,
		},
		{
			name: "invalid HTTPSProxy",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "https://bad%20uri"
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "https://bad%20uri": parse "https://bad%20uri": invalid URL escape "%20"$`,
		},
		{
			name: "invalid HTTPSProxy Schema missing",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "http//baduri"
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Invalid value: "http//baduri": parse "http//baduri": invalid URI for request$`,
		},
		{
			name: "invalid HTTPSProxy Schema different schema",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.HTTPSProxy = "ftp://baduri"
				return c
			}(),
			expectedError: `^proxy.httpsProxy: Unsupported value: "ftp": supported values: "http", "https"$`,
		},
		{
			name: "invalid NoProxy domain",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.NoProxy = "good-no-proxy.com,*.bad-proxy"
				return c
			}(),
			expectedError: `^\Qproxy.noProxy: Invalid value: "good-no-proxy.com,*.bad-proxy": each element of noProxy must be a IP, CIDR or domain without wildcard characters, which is violated by element 1 "*.bad-proxy"\E$`,
		},
		{
			name: "invalid NoProxy spaces",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.NoProxy = "good-no-proxy.com, *.bad-proxy"
				return c
			}(),
			expectedError: `^\Q[proxy.noProxy: Invalid value: "good-no-proxy.com, *.bad-proxy": noProxy must not have spaces, proxy.noProxy: Invalid value: "good-no-proxy.com, *.bad-proxy": each element of noProxy must be a IP, CIDR or domain without wildcard characters, which is violated by element 1 "*.bad-proxy"]\E$`,
		},
		{
			name: "invalid NoProxy CIDR",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.NoProxy = "good-no-proxy.com,172.bad.CIDR.0/16"
				return c
			}(),
			expectedError: `^\Qproxy.noProxy: Invalid value: "good-no-proxy.com,172.bad.CIDR.0/16": each element of noProxy must be a IP, CIDR or domain without wildcard characters, which is violated by element 1 "172.bad.CIDR.0/16"\E$`,
		},
		{
			name: "invalid NoProxy domain & CIDR",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.NoProxy = "good-no-proxy.com,a-good-one,*.bad-proxy.,another,172.bad.CIDR.0/16,good-end"
				return c
			}(),
			expectedError: `^\Q[proxy.noProxy: Invalid value: "good-no-proxy.com,a-good-one,*.bad-proxy.,another,172.bad.CIDR.0/16,good-end": each element of noProxy must be a IP, CIDR or domain without wildcard characters, which is violated by element 2 "*.bad-proxy.", proxy.noProxy: Invalid value: "good-no-proxy.com,a-good-one,*.bad-proxy.,another,172.bad.CIDR.0/16,good-end": each element of noProxy must be a IP, CIDR or domain without wildcard characters, which is violated by element 4 "172.bad.CIDR.0/16"]\E$`,
		},
		{
			name: "valid * NoProxy",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Proxy.NoProxy = "*"
				return c
			}(),
		},
		{
			name: "release image source is not valid",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source: "ocp/release-x.y",
				}}
				return c
			}(),
			expectedError: `^imageContentSources\[0\]\.source: Invalid value: "ocp/release-x\.y": the repository provided is invalid: a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, \'\-\' or \'\.\', and must start and end with an alphanumeric character \(e.g. \'example\.com\', regex used for validation is \'\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\(\\\.\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\)\*\'\)`,
		},
		{
			name: "release image source's mirror is not valid",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source:  "q.io/ocp/release-x.y",
					Mirrors: []string{"ocp/openshift-x.y"},
				}}
				return c
			}(),
			expectedError: `^imageContentSources\[0\]\.mirrors\[0\]: Invalid value: "ocp/openshift-x\.y": the repository provided is invalid: a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, \'\-\' or \'\.\', and must start and end with an alphanumeric character \(e.g. \'example\.com\', regex used for validation is \'\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\(\\\.\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\)\*\'\)`,
		},
		{
			name: "release image source's mirror is valid",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source:  "q.io/ocp/release-x.y",
					Mirrors: []string{"mirror.example.com:5000"},
				}}
				return c
			}(),
		},
		{
			name: "release image source is not repository but reference by digest",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source: "quay.io/ocp/release-x.y@sha256:397c867cc10bcc90cf05ae9b71dd3de6000535e27cb6c704d9f503879202582c",
				}}
				return c
			}(),
			expectedError: `^imageContentSources\[0\]\.source: Invalid value: "quay\.io/ocp/release-x\.y@sha256:397c867cc10bcc90cf05ae9b71dd3de6000535e27cb6c704d9f503879202582c": must be repository--not reference$`,
		},
		{
			name: "release image source is not repository but reference by tag",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source: "quay.io/ocp/release-x.y:latest",
				}}
				return c
			}(),
			expectedError: `^imageContentSources\[0\]\.source: Invalid value: "quay\.io/ocp/release-x\.y:latest": must be repository--not reference$`,
		},
		{
			name: "valid release image source",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source: "quay.io/ocp/release-x.y",
				}}
				return c
			}(),
		},
		{
			name: "release image source is not valid ImageDigestSource",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source: "ocp/release-x.y",
				}}
				return c
			}(),
			expectedError: `^imageDigestSources\[0\]\.source: Invalid value: "ocp/release-x\.y": the repository provided is invalid: a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, \'\-\' or \'\.\', and must start and end with an alphanumeric character \(e.g. \'example\.com\', regex used for validation is \'\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\(\\\.\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\)\*\'\)`,
		},
		{
			name: "release image source's mirror is not valid ImageDigestSource",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source:  "q.io/ocp/release-x.y",
					Mirrors: []string{"ocp/openshift-x.y"},
				}}
				return c
			}(),
			expectedError: `^imageDigestSources\[0\]\.mirrors\[0\]: Invalid value: "ocp/openshift-x\.y": the repository provided is invalid: a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, \'\-\' or \'\.\', and must start and end with an alphanumeric character \(e.g. \'example\.com\', regex used for validation is \'\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\(\\\.\[a\-z0\-9\]\(\[\-a\-z0\-9\]\*\[a\-z0\-9\]\)\?\)\*\'\)`,
		},
		{
			name: "release image source's mirror is valid ImageDigestSource",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source:  "q.io/ocp/release-x.y",
					Mirrors: []string{"mirror.example.com:5000"},
				}}
				return c
			}(),
		},
		{
			name: "release image source is not repository but reference by digest ImageDigestSource",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source: "quay.io/ocp/release-x.y@sha256:397c867cc10bcc90cf05ae9b71dd3de6000535e27cb6c704d9f503879202582c",
				}}
				return c
			}(),
			expectedError: `^imageDigestSources\[0\]\.source: Invalid value: "quay\.io/ocp/release-x\.y@sha256:397c867cc10bcc90cf05ae9b71dd3de6000535e27cb6c704d9f503879202582c": must be repository--not reference$`,
		},
		{
			name: "release image source is not repository but reference by tag ImageDigestSource",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source: "quay.io/ocp/release-x.y:latest",
				}}
				return c
			}(),
			expectedError: `^imageDigestSources\[0\]\.source: Invalid value: "quay\.io/ocp/release-x\.y:latest": must be repository--not reference$`,
		},
		{
			name: "valid release image source ImageDigstSourrce",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source: "quay.io/ocp/release-x.y",
				}}
				return c
			}(),
		},
		{
			name: "error out ImageContentSources and ImageDigestSources and are set at the same time",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.DeprecatedImageContentSources = []types.ImageContentSource{{
					Source:  "q.io/ocp/source",
					Mirrors: []string{"ocp/openshift/mirror"},
				}}
				c.ImageDigestSources = []types.ImageDigestSource{{
					Source:  "q.io/ocp/source",
					Mirrors: []string{"ocp-digest/openshift/mirror"}}}
				return c
			}(),
			expectedError: `cannot set imageContentSources and imageDigestSources at the same time`,
		},
		{
			name: "invalid publishing strategy",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Publish = types.PublishingStrategy("ExternalInternalDoNotCare")
				return c
			}(),
			expectedError: `^publish: Unsupported value: \"ExternalInternalDoNotCare\": supported values: \"External\", \"Internal\"`,
		},
		{
			name: "architecture is not supported",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute[0].Architecture = types.ArchitectureS390X
				c.ControlPlane.Architecture = types.ArchitectureS390X
				return c
			}(),
			expectedError: `[controlPlane.architecture: Unsupported value: "s390x": supported values: "amd64", "arm64", compute\[0\].architecture: Unsupported value: "s390x": supported values: "amd64", "arm64"]`,
		},
		{
			name: "architecture is not supported",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute[0].Architecture = types.ArchitecturePPC64LE
				c.ControlPlane.Architecture = types.ArchitecturePPC64LE
				return c
			}(),
			expectedError: `[controlPlane.architecture: Unsupported value: "ppc64le": supported values: "amd64", "arm64", compute\[0\].architecture: Unsupported value: "ppc64le": supported values: "amd64", "arm64"]`,
		},
		{
			name: "cluster is not heteregenous",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Compute[0].Architecture = types.ArchitectureARM64
				return c
			}(),
			expectedError: `^compute\[0\].architecture: Invalid value: "arm64": heteregeneous multi-arch is not supported; compute pool architecture must match control plane$`,
		},
		{
			name: "valid cloud credentials mode",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.CredentialsMode = types.PassthroughCredentialsMode
				return c
			}(),
		},
		{
			name: "bad cloud credentials mode",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.CredentialsMode = "bad-mode"
				return c
			}(),
			expectedError: `^credentialsMode: Unsupported value: "bad-mode": supported values: "Manual", "Mint", "Passthrough"$`,
		},
		{
			name: "allowed docker bridge with non-libvirt",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Networking.MachineNetwork = []types.MachineNetworkEntry{{CIDR: *ipnet.MustParseCIDR("172.17.64.0/18")}}
				return c
			}(),
			expectedError: ``,
		},
		{
			name: "valid baseline capability set",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: "v4.11"}
				return c
			}(),
		},
		{
			name: "invalid empty string baseline capability set",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: ""}
				return c
			}(),
			expectedError: `capabilities.baselineCapabilitySet: Unsupported value: "": supported values: .*`,
		},
		{
			name: "invalid baseline capability set specified",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: "vNotValid"}
				return c
			}(),
			expectedError: `capabilities.baselineCapabilitySet: Unsupported value: "vNotValid": supported values: .*`,
		},
		{
			name: "valid additional enabled capability specified",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: "v4.11",
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{"openshift-samples"}}
				return c
			}(),
		},
		{
			name: "invalid empty additional enabled capability specified",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: "v4.11",
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{""}}
				return c
			}(),
			expectedError: `capabilities.additionalEnabledCapabilities\[0\]: Unsupported value: "": supported values: .*`,
		},
		{
			name: "invalid additional enabled capability specified",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{BaselineCapabilitySet: "v4.11",
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{"not-valid"}}
				return c
			}(),
			expectedError: `capabilities.additionalEnabledCapabilities\[0\]: Unsupported value: "not-valid": supported values: .*`,
		},
		//VIP tests

		{
			name: "should validate vips on OpenStack (vips are required on openstack)",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Platform = types.Platform{
					OpenStack: validOpenStackPlatform(),
				}
				c.Platform.OpenStack.DeprecatedAPIVIP = ""
				c.Platform.OpenStack.APIVIPs = []string{}

				return c
			}(),
			expectedError: "platform.openstack.apiVIPs: Required value: must specify at least one VIP for the API",
		},
		// {
		// 	name: "should not validate vips on OpenStack if not set (vips are not required on openstack)",
		// 	installConfig: func() *types.InstallConfig {
		// 		c := validInstallConfig()
		// 		c.Platform = types.Platform{
		// 			OpenStack: validOpenStackPlatform(),
		// 		}
		// 		c.Platform.OpenStack.DeprecatedAPIVIP = ""
		// 		c.Platform.OpenStack.APIVIPs = []string{}

		// 		return c
		// 	}(),
		// },
		{
			name: "should validate vips on OpenStack if set (vips are not required on openstack)",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Platform = types.Platform{
					OpenStack: validOpenStackPlatform(),
				}
				c.Platform.OpenStack.APIVIPs = []string{"foobar"}

				return c
			}(),
			expectedError: "platform.openstack.apiVIPs: Invalid value: \"foobar\": \"foobar\" is not a valid IP",
		},
		{
			name: "valid custom features",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = configv1.CustomNoUpgrade
				c.FeatureGates = []string{
					"CustomFeature1=True",
					"CustomFeature2=False",
				}
				return c
			}(),
		},
		{
			name: "invalid custom features",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = configv1.CustomNoUpgrade
				c.FeatureGates = []string{
					"CustomFeature1=True",
					"CustomFeature2",
				}
				return c
			}(),
			expectedError: `featureGates\[1\]: Invalid value: "CustomFeature2": must match the format <feature-name>=<bool>`,
		},
		{
			name: "invalid custom features bool",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = configv1.CustomNoUpgrade
				c.FeatureGates = []string{
					"CustomFeature1=foo",
					"CustomFeature2=False",
				}
				return c
			}(),
			expectedError: `featureGates\[0\]: Invalid value: "CustomFeature1=foo": must match the format <feature-name>=<bool>, could not parse boolean value`,
		},
		{
			name: "custom features supplied with non-custom featureset",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.FeatureSet = configv1.TechPreviewNoUpgrade
				c.FeatureGates = []string{
					"CustomFeature1=True",
					"CustomFeature2=False",
				}
				return c
			}(),
			expectedError: "featureGates: Forbidden: featureGates can only be used with the CustomNoUpgrade feature set",
		},
		{
			name: "return error when MAPI disabled w/o baremetal with baseline none",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{
					BaselineCapabilitySet:         configv1.ClusterVersionCapabilitySetNone,
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{configv1.ClusterVersionCapabilityBaremetal},
				}
				return c
			}(),
			expectedError: `the baremetal capability requires the MachineAPI capability`,
		},
		{
			name: "valid disabled MAPI capability configuration",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{
					BaselineCapabilitySet: configv1.ClusterVersionCapabilitySetNone,
				}
				return c
			}(),
		},
		{
			name: "valid enabled MAPI capability configuration",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{
					BaselineCapabilitySet:         configv1.ClusterVersionCapabilitySetNone,
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{configv1.ClusterVersionCapabilityBaremetal, configv1.ClusterVersionCapabilityMachineAPI},
				}
				return c
			}(),
		},
		{
			name: "valid enabled MAPI capability configuration 2",
			installConfig: func() *types.InstallConfig {
				c := validInstallConfig()
				c.Capabilities = &types.Capabilities{
					BaselineCapabilitySet:         configv1.ClusterVersionCapabilitySetNone,
					AdditionalEnabledCapabilities: []configv1.ClusterVersionCapability{configv1.ClusterVersionCapabilityMachineAPI},
				}
				return c
			}(),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateInstallConfig(tc.installConfig, false).ToAggregate()
			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Regexp(t, tc.expectedError, err)
			}
		})
	}
}

func Test_ensureIPv4IsFirstInDualStackSlice(t *testing.T) {
	tests := []struct {
		name    string
		vips    []string
		want    []string
		wantErr bool
	}{
		{
			name:    "should switch VIPs",
			vips:    []string{"fe80::0", "192.168.1.1"},
			want:    []string{"192.168.1.1", "fe80::0"},
			wantErr: false,
		},
		{
			name:    "should do nothing on single stack",
			vips:    []string{"192.168.1.1"},
			want:    []string{"192.168.1.1"},
			wantErr: false,
		},
		{
			name:    "should do nothing on correct order",
			vips:    []string{"192.168.1.1", "fe80::0"},
			want:    []string{"192.168.1.1", "fe80::0"},
			wantErr: false,
		},
		{
			name:    "return error on invalid number of vips",
			vips:    []string{"192.168.1.1", "fe80::0", "192.168.1.1"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ensureIPv4IsFirstInDualStackSlice(&tt.vips, field.NewPath("test")); (len(err) > 0) != tt.wantErr {
				t.Errorf("ensureIPv4IsFirstInDualStackSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !utilsslice.Equal(tt.vips, tt.want) && len(tt.vips) == 2 {
				t.Errorf("ensureIPv4IsFirstInDualStackSlice() changed to %v, expected %v", tt.vips, tt.want)
			}
		})
	}
}
