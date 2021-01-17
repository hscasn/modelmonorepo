package versionsets

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// PlatformVersionset is a set of versions for a specific platform
type PlatformVersionset map[string]string

// Versionset is a versionset file
type Versionset struct {
	CloudRun PlatformVersionset `yaml:"cloudrun"`
}

// Parse yaml versionset
func Parse(yamlVersionset []byte) (Versionset, error) {
	t := Versionset{}

	err := yaml.Unmarshal(yamlVersionset, &t)
	if err != nil {
		return t, fmt.Errorf("error parsing versionset: %w", err)
	}

	return t, nil
}
