package conversion

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation/field"
	utilsslice "k8s.io/utils/strings/slices"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	operv1 "github.com/openshift/api/operator/v1"
)

// ConvertInstallConfig is modeled after the k8s conversion schemes, which is
// how deprecated values are upconverted.
// This updates the APIVersion to reflect the fact that we've internally
// upconverted.
func ConvertInstallConfig(config *types.InstallConfig) error {
	// check that the version is convertible
	switch config.APIVersion {
	case types.InstallConfigVersion, "v1beta3", "v1beta4":
		// works
	case "":
		return field.Required(field.NewPath("apiVersion"), "no version was provided")
	default:
		return field.Invalid(field.NewPath("apiVersion"), config.APIVersion, fmt.Sprintf("cannot upconvert from version %s", config.APIVersion))
	}
	convertNetworking(config)
	switch config.Platform.Name() {
	case openstack.Name:
		if err := convertOpenStack(config); err != nil {
			return err
		}
	}

	config.APIVersion = types.InstallConfigVersion
	return nil
}

// convertNetworking upconverts deprecated fields in networking
func convertNetworking(config *types.InstallConfig) {
	if config.Networking == nil {
		return
	}

	netconf := config.Networking

	if len(netconf.ClusterNetwork) == 0 {
		netconf.ClusterNetwork = netconf.DeprecatedClusterNetworks
	}

	if len(netconf.MachineNetwork) == 0 && netconf.DeprecatedMachineCIDR != nil {
		netconf.MachineNetwork = []types.MachineNetworkEntry{
			{CIDR: *netconf.DeprecatedMachineCIDR},
		}
	}

	if len(netconf.ServiceNetwork) == 0 && netconf.DeprecatedServiceCIDR != nil {
		netconf.ServiceNetwork = []ipnet.IPNet{*netconf.DeprecatedServiceCIDR}
	}

	// Convert type to networkType if the latter is missing
	if netconf.NetworkType == "" {
		netconf.NetworkType = netconf.DeprecatedType
	}

	// Recognize the OpenShiftSDN network plugin name regardless of capitalization, for
	// backward compatibility
	if strings.ToLower(netconf.NetworkType) == strings.ToLower(string(operv1.NetworkTypeOpenShiftSDN)) {
		netconf.NetworkType = string(operv1.NetworkTypeOpenShiftSDN)
	}

	// Convert hostSubnetLength to hostPrefix
	for i, entry := range netconf.ClusterNetwork {
		if entry.HostPrefix == 0 && entry.DeprecatedHostSubnetLength != 0 {
			_, size := entry.CIDR.Mask.Size()
			netconf.ClusterNetwork[i].HostPrefix = int32(size) - entry.DeprecatedHostSubnetLength
		}
	}
}

// convertOpenStack upconverts deprecated fields in the OpenStack platform.
func convertOpenStack(config *types.InstallConfig) error {
	// LbFloatingIP has been renamed to APIFloatingIP
	if config.Platform.OpenStack.DeprecatedLbFloatingIP != "" {
		if config.Platform.OpenStack.APIFloatingIP == "" {
			config.Platform.OpenStack.APIFloatingIP = config.Platform.OpenStack.DeprecatedLbFloatingIP
		} else if config.Platform.OpenStack.DeprecatedLbFloatingIP != config.Platform.OpenStack.APIFloatingIP {
			// Return error if both LbFloatingIP and APIFloatingIP are specified in the config
			return field.Forbidden(field.NewPath("platform").Child("openstack").Child("lbFloatingIP"), "cannot specify lbFloatingIP and apiFloatingIP together")
		}
	}

	// computeFlavor has been deprecated in favor of type in defaultMachinePlatform.
	if config.Platform.OpenStack.DeprecatedFlavorName != "" {
		if config.Platform.OpenStack.DefaultMachinePlatform == nil {
			config.Platform.OpenStack.DefaultMachinePlatform = &openstack.MachinePool{}
		}

		if config.Platform.OpenStack.DefaultMachinePlatform.FlavorName != "" && config.Platform.OpenStack.DefaultMachinePlatform.FlavorName != config.Platform.OpenStack.DeprecatedFlavorName {
			// Return error if both computeFlavor and type of defaultMachinePlatform are specified in the config
			return field.Forbidden(field.NewPath("platform").Child("openstack").Child("computeFlavor"), "cannot specify computeFlavor and type in defaultMachinePlatform together")
		}

		config.Platform.OpenStack.DefaultMachinePlatform.FlavorName = config.Platform.OpenStack.DeprecatedFlavorName
	}

	// type has been deprecated in favor of types in the machinePools.
	if config.ControlPlane != nil &&
		config.ControlPlane.Platform.OpenStack != nil &&
		config.ControlPlane.Platform.OpenStack.RootVolume != nil &&
		config.ControlPlane.Platform.OpenStack.RootVolume.DeprecatedType != "" {
		if len(config.ControlPlane.Platform.OpenStack.RootVolume.Types) > 0 {
			// Return error if both type and types of rootVolume are specified in the config
			return field.Forbidden(field.NewPath("controlPlane").Child("platform").Child("openstack").Child("rootVolume").Child("type"), "cannot specify type and types in rootVolume together")
		}
		config.ControlPlane.Platform.OpenStack.RootVolume.Types = []string{config.ControlPlane.Platform.OpenStack.RootVolume.DeprecatedType}
		config.ControlPlane.Platform.OpenStack.RootVolume.DeprecatedType = ""
	}
	for _, pool := range config.Compute {
		mpool := pool.Platform.OpenStack
		if mpool != nil && mpool.RootVolume != nil && mpool.RootVolume.DeprecatedType != "" {
			if mpool.RootVolume.Types != nil && len(mpool.RootVolume.Types) > 0 {
				// Return error if both type and types of rootVolume are specified in the config
				return field.Forbidden(field.NewPath("compute").Child("platform").Child("openstack").Child("rootVolume").Child("type"), "cannot specify type and types in rootVolume together")
			}
			mpool.RootVolume.Types = []string{mpool.RootVolume.DeprecatedType}
			mpool.RootVolume.DeprecatedType = ""
		}
	}
	if config.Platform.OpenStack.DefaultMachinePlatform != nil && config.Platform.OpenStack.DefaultMachinePlatform.RootVolume != nil && config.Platform.OpenStack.DefaultMachinePlatform.RootVolume.DeprecatedType != "" {
		if len(config.Platform.OpenStack.DefaultMachinePlatform.RootVolume.Types) > 0 {
			// Return error if both type and types of defaultMachinePlatform are specified in the config
			return field.Forbidden(field.NewPath("platform").Child("openstack").Child("type"), "cannot specify type and types in defaultMachinePlatform together")
		}
		config.Platform.OpenStack.DefaultMachinePlatform.RootVolume.Types = []string{config.Platform.OpenStack.DefaultMachinePlatform.RootVolume.DeprecatedType}
		config.Platform.OpenStack.DefaultMachinePlatform.RootVolume.DeprecatedType = ""
	}

	if err := upconvertVIP(&config.Platform.OpenStack.APIVIPs, config.Platform.OpenStack.DeprecatedAPIVIP, "apiVIP", "apiVIPs", field.NewPath("platform").Child("openstack")); err != nil {
		return err
	}

	if err := upconvertVIP(&config.Platform.OpenStack.IngressVIPs, config.Platform.OpenStack.DeprecatedIngressVIP, "ingressVIP", "ingressVIPs", field.NewPath("platform").Child("openstack")); err != nil {
		return err
	}

	// machinesSubnet has been deprecated in favor of ControlPlanePort
	controlPlanePort := config.Platform.OpenStack.ControlPlanePort
	deprecatedMachinesSubnet := config.Platform.OpenStack.DeprecatedMachinesSubnet
	if deprecatedMachinesSubnet != "" && controlPlanePort == nil {
		fixedIPs := []openstack.FixedIP{{Subnet: openstack.SubnetFilter{ID: deprecatedMachinesSubnet}}}
		config.Platform.OpenStack.ControlPlanePort = &openstack.PortTarget{FixedIPs: fixedIPs}
	} else if deprecatedMachinesSubnet != "" &&
		controlPlanePort != nil {
		if !(len(controlPlanePort.FixedIPs) == 1 && controlPlanePort.FixedIPs[0].Subnet.ID == deprecatedMachinesSubnet) {
			return field.Invalid(field.NewPath("platform").Child("openstack").Child("machinesSubnet"), deprecatedMachinesSubnet, fmt.Sprintf("%s is deprecated; only %s needs to be specified", "machinesSubnet", "controlPlanePort"))
		}
	}

	return nil
}

// upconvertVIP upconverts the deprecated VIP (oldVIPValue) to the new VIPs
// slice (newVIPValues). It returns errors, if both fields are set and all
// contain unique values
func upconvertVIP(newVIPValues *[]string, oldVIPValue, newFieldName, oldFieldName string, fldPath *field.Path) error {
	if oldVIPValue != "" && len(*newVIPValues) == 0 {
		*newVIPValues = []string{oldVIPValue}
	} else if oldVIPValue != "" &&
		len(*newVIPValues) > 0 &&
		!utilsslice.Contains(*newVIPValues, oldVIPValue) {

		return field.Invalid(fldPath.Child(oldFieldName), oldVIPValue, fmt.Sprintf("%s is deprecated; only %s needs to be specified", oldFieldName, newFieldName))
	}

	return nil
}
