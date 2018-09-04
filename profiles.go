// Copyright 2011, 2012, 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package charm

import (
	"fmt"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Profiles struct {
	Providers map[string]interface{} `json:"providers" yaml:"providers"`
}

func NewProfiles() *Profiles {
	return &Profiles{}
}

// ReadProfiles reads in a Profile from a charm's profile.yaml.
// It's not validated at this point, so that the caller can choose to
// override any validation.
// The expected format of the yaml is the following:
//
// ```
// providers:
//   - name: lxd
//     config.value: "true"
//     devices:
//       gpu:
//       type: gpu
// ```
func ReadProfiles(r io.Reader) (*Profiles, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var profile *Profiles
	if err := yaml.Unmarshal(data, &profile); err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, fmt.Errorf("invalid profile.yaml: empty profile")
	}
	return profile, nil
}
