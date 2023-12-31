package installconfig

import (
	"fmt"
	"sort"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/pkg/errors"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	openstackconfig "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// Platform is an asset that queries the user for the platform on which to install
// the cluster.
type platform struct {
	types.Platform
}

var _ asset.Asset = (*platform)(nil)

// Dependencies returns no dependencies.
func (a *platform) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate queries for input from the user.
func (a *platform) Generate(asset.Parents) error {
	platform, err := a.queryUserForPlatform()
	if err != nil {
		return err
	}

	switch platform {
	case openstack.Name:
		a.OpenStack, err = openstackconfig.Platform()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown platform type %q", platform)
	}

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *platform) Name() string {
	return "Platform"
}

func (a *platform) queryUserForPlatform() (platform string, err error) {
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Platform",
				Options: types.PlatformNames,
				Help:    "The platform on which the cluster will run.  For a full list of platforms, including those not supported by this wizard, see https://github.com/anton-sidelnikov/otc-openshift-installer",
			},
			Validate: survey.ComposeValidators(survey.Required, func(ans interface{}) error {
				choice := ans.(core.OptionAnswer).Value
				i := sort.SearchStrings(types.PlatformNames, choice)
				if i == len(types.PlatformNames) || types.PlatformNames[i] != choice {
					return errors.Errorf("invalid platform %q", choice)
				}
				return nil
			}),
		},
	}, &platform)
	return
}

func (a *platform) CurrentName() string {
	return a.Platform.Name()
}
