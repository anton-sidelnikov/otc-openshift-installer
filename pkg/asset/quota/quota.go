package quota

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	openstackvalidation "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig/openstack/validation"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/quota/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/diagnostics"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/quota"
	typesopenstack "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// PlatformQuotaCheck is an asset that validates the install-config platform for
// any resource requirements based on the quotas available.
type PlatformQuotaCheck struct {
}

var _ asset.Asset = (*PlatformQuotaCheck)(nil)

// Dependencies returns the dependencies for PlatformQuotaCheck
func (a *PlatformQuotaCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
		&machines.Master{},
		&machines.Worker{},
	}
}

// Generate queries for input from the user.
func (a *PlatformQuotaCheck) Generate(dependencies asset.Parents) error {
	ic := &installconfig.InstallConfig{}
	mastersAsset := &machines.Master{}
	workersAsset := &machines.Worker{}
	dependencies.Get(ic, mastersAsset, workersAsset)

	masters, err := mastersAsset.Machines()
	if err != nil {
		return err
	}

	workers, err := workersAsset.MachineSets()
	if err != nil {
		return err
	}

	platform := ic.Config.Platform.Name()
	switch platform {
	case typesopenstack.Name:
		if skip := os.Getenv("OPENSHIFT_INSTALL_SKIP_PREFLIGHT_VALIDATIONS"); skip == "1" {
			logrus.Warnf("OVERRIDE: pre-flight validation disabled.")
			return nil
		}
		ci, err := openstackvalidation.GetCloudInfo(ic.Config)
		if err != nil {
			return errors.Wrap(err, "failed to get cloud info")
		}
		if ci == nil {
			logrus.Warnf("Empty OpenStack cloud info and therefore will skip checking quota validation.")
			return nil
		}
		reports, err := quota.Check(ci.Quotas, openstack.Constraints(ci, masters, workers, ic.Config.NetworkType))
		if err != nil {
			return summarizeFailingReport(reports)
		}
		summarizeReport(reports)
	default:
		err = fmt.Errorf("unknown platform type %q", platform)
	}
	return err
}

// Name returns the human-friendly name of the asset.
func (a *PlatformQuotaCheck) Name() string {
	return "Platform Quota Check"
}

// summarizeFailingReport summarizes a report when there are failing constraints.
func summarizeFailingReport(reports []quota.ConstraintReport) error {
	var notavailable []string
	var unknown []string
	var regionMessage string
	for _, report := range reports {
		switch report.Result {
		case quota.NotAvailable:
			if report.For.Region != "" {
				regionMessage = " in " + report.For.Region
			} else {
				regionMessage = ""
			}
			notavailable = append(notavailable, fmt.Sprintf("%s is not available%s because %s", report.For.Name, regionMessage, report.Message))
		case quota.Unknown:
			unknown = append(unknown, report.For.Name)
		default:
			continue
		}
	}

	if len(notavailable) == 0 && len(unknown) > 0 {
		// all quotas are missing information so warn and skip
		logrus.Warnf("Failed to find information on quotas %s", strings.Join(unknown, ", "))
		return nil
	}

	msg := strings.Join(notavailable, ", ")
	if len(unknown) > 0 {
		msg = fmt.Sprintf("%s, and could not find information on %s", msg, strings.Join(unknown, ", "))
	}
	return &diagnostics.Err{Reason: "MissingQuota", Message: msg}
}

// summarizeReport summarizes a report when there are availble.
func summarizeReport(reports []quota.ConstraintReport) {
	var low []string
	var regionMessage string
	for _, report := range reports {
		switch report.Result {
		case quota.AvailableButLow:
			if report.For.Region != "" {
				regionMessage = " (" + report.For.Region + ")"
			} else {
				regionMessage = ""
			}
			low = append(low, fmt.Sprintf("%s%s", report.For.Name, regionMessage))
		default:
			continue
		}
	}
	if len(low) > 0 {
		logrus.Warnf("Following quotas %s are available but will be completely used pretty soon.", strings.Join(low, ", "))
	}
}
