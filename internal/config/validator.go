package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// ConfigMetadata represents the required fields for all config files
type ConfigMetadata struct {
	ID          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

// ValidateConfig validates the required fields in a YAML config
func ValidateConfig(data []byte) error {
	var metadata ConfigMetadata

	// First try to unmarshal just the metadata fields
	if err := yaml.Unmarshal(data, &metadata); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	// Validate required fields
	if metadata.ID == "" {
		return fmt.Errorf("missing required field: id")
	}
	if metadata.Title == "" {
		return fmt.Errorf("missing required field: title")
	}
	if metadata.Description == "" {
		return fmt.Errorf("missing required field: description")
	}
	if metadata.Version == "" {
		return fmt.Errorf("missing required field: version")
	}

	return nil
}
