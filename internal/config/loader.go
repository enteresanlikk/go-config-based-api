package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	instance *ConfigLoader
	once     sync.Once
)

type ConfigLoader struct {
	configs map[string]interface{}
	mu      sync.RWMutex
}

// ConfigData represents the structure of config files with ID
type ConfigData struct {
	ID string `yaml:"id"`
}

// GetInstance returns the singleton instance of ConfigLoader
func GetInstance() *ConfigLoader {
	once.Do(func() {
		instance = &ConfigLoader{
			configs: make(map[string]interface{}),
		}
		err := instance.LoadConfigs()
		if err != nil {
			panic(fmt.Sprintf("Failed to load configs: %v", err))
		}
	})
	return instance
}

// LoadConfigs loads all YAML files from the configs directory recursively
func (c *ConfigLoader) LoadConfigs() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.loadConfigsRecursive("configs")
}

// loadConfigsRecursive recursively loads YAML files from directories
func (c *ConfigLoader) loadConfigsRecursive(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())

		if file.IsDir() {
			// Recursively process subdirectories
			if err := c.loadConfigsRecursive(path); err != nil {
				return err
			}
			continue
		}

		// Only process YAML files
		if filepath.Ext(file.Name()) != ".yml" {
			continue
		}

		// Read and parse the config file
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		// Validate config file
		if err := ValidateConfig(data); err != nil {
			return fmt.Errorf("validation failed for file %s: %w", path, err)
		}

		// Now unmarshal the full config
		var config interface{}
		if err := yaml.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("failed to unmarshal file %s: %w", path, err)
		}

		// Get the ID from the validated config
		var metadata ConfigMetadata
		if err := yaml.Unmarshal(data, &metadata); err != nil {
			return fmt.Errorf("failed to get metadata from file %s: %w", path, err)
		}

		c.configs[metadata.ID] = config
	}

	return nil
}

// GetConfig returns the config for the given ID
func (c *ConfigLoader) GetConfig(id string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	config, exists := c.configs[id]
	return config, exists
}

// GetAllConfigs returns all loaded configs
func (c *ConfigLoader) GetAllConfigs() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	configsCopy := make(map[string]interface{})
	for k, v := range c.configs {
		configsCopy[k] = v
	}
	return configsCopy
}
