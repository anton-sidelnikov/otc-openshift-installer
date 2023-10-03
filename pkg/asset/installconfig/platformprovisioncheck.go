package installconfig

import (
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	osconfig "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	configv1 "github.com/openshift/api/config/v1"
)

// PlatformProvisionCheck is an asset that validates the install-config platform for
// any requirements specific for provisioning infrastructure.
type PlatformProvisionCheck struct {
}

var _ asset.Asset = (*PlatformProvisionCheck)(nil)

// Dependencies returns the dependencies for PlatformProvisionCheck
func (a *PlatformProvisionCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&InstallConfig{},
	}
}

// Generate queries for input from the user.
func (a *PlatformProvisionCheck) Generate(dependencies asset.Parents) error {
	ic := &InstallConfig{}
	dependencies.Get(ic)
	platform := ic.Config.Platform.Name()

	// IPI requires MachineAPI capability
	enabledCaps := sets.NewString()
	if ic.Config.Capabilities == nil || ic.Config.Capabilities.BaselineCapabilitySet == "" {
		// when Capabilities and/or BaselineCapabilitySet is not specified, default is vCurrent
		baseSet := configv1.ClusterVersionCapabilitySets[configv1.ClusterVersionCapabilitySetCurrent]
		for _, cap := range baseSet {
			enabledCaps.Insert(string(cap))
		}
	}
	if ic.Config.Capabilities != nil {
		if ic.Config.Capabilities.BaselineCapabilitySet != "" {
			baseSet := configv1.ClusterVersionCapabilitySets[ic.Config.Capabilities.BaselineCapabilitySet]
			for _, cap := range baseSet {
				enabledCaps.Insert(string(cap))
			}
		}
		if ic.Config.Capabilities.AdditionalEnabledCapabilities != nil {
			for _, cap := range ic.Config.Capabilities.AdditionalEnabledCapabilities {
				enabledCaps.Insert(string(cap))
			}
		}
	}
	if !enabledCaps.Has(string(configv1.ClusterVersionCapabilityMachineAPI)) {
		return errors.New("IPI requires MachineAPI capability")
	}

	switch platform {
	case openstack.Name:
		err := osconfig.ValidateForProvisioning(ic.Config)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown platform type %q", platform)
	}
	return nil
}

// Name returns the human-friendly name of the asset.
func (a *PlatformProvisionCheck) Name() string {
	return "Platform Provisioning Check"
}
