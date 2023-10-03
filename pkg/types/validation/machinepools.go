package validation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	openstackvalidation "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack/validation"
)

var (
	validHyperthreadingModes = map[types.HyperthreadingMode]bool{
		types.HyperthreadingDisabled: true,
		types.HyperthreadingEnabled:  true,
	}

	validHyperthreadingModeValues = func() []string {
		v := make([]string, 0, len(validHyperthreadingModes))
		for m := range validHyperthreadingModes {
			v = append(v, string(m))
		}
		return v
	}()

	validArchitectures = map[types.Architecture]bool{
		types.ArchitectureAMD64:   true,
		types.ArchitectureS390X:   true,
		types.ArchitecturePPC64LE: true,
		types.ArchitectureARM64:   true,
	}

	validArchitectureValues = func() []string {
		v := make([]string, 0, len(validArchitectures))
		for m := range validArchitectures {
			v = append(v, string(m))
		}
		return v
	}()
)

// ValidateMachinePool checks that the specified machine pool is valid.
func ValidateMachinePool(platform *types.Platform, p *types.MachinePool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if p.Replicas != nil {
		if *p.Replicas < 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("replicas"), p.Replicas, "number of replicas must not be negative"))
		}
	} else {
		allErrs = append(allErrs, field.Required(fldPath.Child("replicas"), "replicas is required"))
	}
	if !validHyperthreadingModes[p.Hyperthreading] {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("hyperthreading"), p.Hyperthreading, validHyperthreadingModeValues))
	}
	if !validArchitectures[p.Architecture] {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("architecture"), p.Architecture, validArchitectureValues))
	}
	allErrs = append(allErrs, validateMachinePoolPlatform(platform, &p.Platform, p, fldPath.Child("platform"))...)
	return allErrs
}

func validateMachinePoolPlatform(platform *types.Platform, p *types.MachinePoolPlatform, pool *types.MachinePool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	platformName := platform.Name()
	validate := func(n string, value interface{}, validation func(*field.Path) field.ErrorList) {
		f := fldPath.Child(n)
		if platformName == n {
			allErrs = append(allErrs, validation(f)...)
		} else {
			allErrs = append(allErrs, field.Invalid(f, value, fmt.Sprintf("cannot specify %q for machine pool when cluster is using %q", n, platformName)))
		}
	}
	if p.OpenStack != nil {
		validate(openstack.Name, p.OpenStack, func(f *field.Path) field.ErrorList {
			return openstackvalidation.ValidateMachinePool(platform.OpenStack, p.OpenStack, pool.Name, f)
		})
	}
	return allErrs
}
