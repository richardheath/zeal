package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Package Zeal package configuration.
type Package struct {
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Version      string              `json:"version"`
	Keep         string              `json:"keep"`
	Split        bool                `json:"split"`
	Author       string              `json:"author"`
	Destination  string              `json:"destination"`
	Files        map[string]string   `json:"files"`
	Exclude      []string            `json:"exclude"`
	Scripts      map[string][]string `json:"scripts"`
	Dependencies map[string]string   `json:"dependencies"`
	Settings     []PackageSetting    `json:"settings"`
	Metadata     map[string]string   `json:"metadata"`
	Override     map[string]Package  `json:"override"`
}

// PackageSetting Zeal package setting configuration.
type PackageSetting struct {
	Default     string `json:"default"`
	Description string `json:"description"`
}

// Load Load zeal config file.
func (config *Package) Load(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	rawConfig, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	json.Unmarshal(rawConfig, config)
	return nil
}

// LoadOverride Populate fields using base config and override config.
func (config *Package) LoadOverride(base Package, targetOverride string) error {
	return nil
}

// ValidateNoSettings Validate that all settings have been replaced
// on specified fiel names.
func (config *Package) ValidateNoSettings(fields []string) error {
	errorMessage := ""
	for _, field := range fields {
		err := config.validateField(field)
		if err != nil {
			errorMessage += err.Error() + "\n"
		}
	}

	if errorMessage != "" {
		return fmt.Errorf(errorMessage)
	}
	return nil
}

// ReplaceSettings Replace specified settings with actual values.
func (config *Package) ReplaceSettings(settings map[string]string) {
}

func (config *Package) replaceSetting(setting PackageSetting, value string) {
}

func replaceString(key string, value string) {

}

func (config *Package) validateField(field string) error {
	switch field {
	case "name":
		return config.checkNoSettingsInString(field, config.Name)
	case "description":
		return config.checkNoSettingsInString(field, config.Description)
	case "version":
		return config.checkNoSettingsInString(field, config.Version)
	case "keep":
		return config.checkNoSettingsInString(field, config.Keep)
	case "author":
		return config.checkNoSettingsInString(field, config.Author)
	case "destination":
		return config.checkNoSettingsInString(field, config.Destination)
	case "files":
		return config.checkNoSettingsInMapString(field, config.Files)
	case "exclude":
		return config.checkNoSettingsInArrayString(field, config.Exclude)
	case "scripts":
		return config.checkNoSettingsInScripts(field, config.Scripts)
	case "dependencies":
		return config.checkNoSettingsInMapString(field, config.Dependencies)
	case "settings":
		return config.checkNoSettingsInSettings(field, config.Settings)
	case "metadata":
		return config.checkNoSettingsInMapString(field, config.Metadata)
	}
	return nil
}

func (config *Package) checkNoSettingsInString(field string, value string) error {
	return nil
}

func (config *Package) checkNoSettingsInMapString(field string, value map[string]string) error {
	return nil
}

func (config *Package) checkNoSettingsInArrayString(field string, value []string) error {
	return nil
}

func (config *Package) checkNoSettingsInScripts(field string, value map[string][]string) error {
	return nil
}

func (config *Package) checkNoSettingsInSettings(field string, value []PackageSetting) error {
	return nil
}
