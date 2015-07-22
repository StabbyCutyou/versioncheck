package versioncheck

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

type registryEntry struct {
	libName    string
	libVer     *version.Version
	constraint *version.Constraint
}

func (r *registryEntry) check() (bool, error) {
	result := r.constraint.Check(r.libVer)
	if result {
		return result, nil
	}
	return result, fmt.Errorf("%s version %s, but required version is %s", r.libName, r.libVer, r.constraint.String())
}

var registry = make([]*registryEntry, 0)

// Register stores information about versions to check, which will apply
// once you call Check()
func Register(name string, libVersion string, constraint string) error {
	libVer, err := version.NewVersion(libVersion)
	if err != nil {
		return err
	}

	lockVer, err := version.NewConstraint(constraint)
	if err != nil {
		return err
	}
	reg := &registryEntry{
		libVer:     libVer,
		constraint: lockVer[0],
		libName:    name,
	}
	registry = append(registry, reg)
	return nil
}

// Check is meant to be called once you have finished Registering your versions.
// Typically, you would run this as a pre-boot on your application or service, and
// either halt or warn on the version mismatch to be alerted to potential changes
// in an imported packages behavior
func Check() []error {
	var errs []error
	for i := 0; i < len(registry); i++ {
		r := registry[i]
		result, err := r.check()
		if !result {
			errs = append(errs, err)
		}
	}
	return errs
}
