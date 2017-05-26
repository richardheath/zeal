package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// Package Zeal package configuration.
type Package struct {
	Name         string                       `json:"name"`
	Description  string                       `json:"description"`
	Version      string                       `json:"version"`
	Keep         string                       `json:"keep"`
	Split        bool                         `json:"split"`
	Author       string                       `json:"author"`
	Destination  string                       `json:"destination"`
	Files        map[string]string            `json:"files"`
	Exclude      []string                     `json:"exclude"`
	Scripts      map[string][]string          `json:"scripts"`
	Dependencies map[string]map[string]string `json:"dependencies"`
	Settings     map[string]PackageSetting    `json:"settings"`
	Metadata     map[string]string            `json:"metadata"`
	Override     map[string]Package           `json:"override"`
}

// PackageSetting Zeal package setting configuration.
type PackageSetting struct {
	Default     string `json:"default"`
	AllowEmpty  bool   `json:"allowEmpty"`
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
	return fmt.Errorf("Not implemented.")
}

// ReplaceTokens Replace specified settings with actual values.
func (config *Package) ReplaceTokens(settings map[string]string) {
	for key, value := range settings {
		fullKey := "{" + key + "}"
		replaceTokens(reflect.ValueOf(config).Elem(), fullKey, value)
	}
}

// ReplaceWithDefaultTokens Replace tokens with default settings for given fields.
func (config *Package) ReplaceWithDefaultTokens(fields []string) {
	for name, setting := range config.Settings {
		for _, field := range fields {
			if setting.Default == "" && setting.AllowEmpty == false {
				// Do not replace tokens where empty is not allowed.
				continue
			}

			key := "#" + name
			config.replaceFieldToken(field, key, setting.Default)
		}
	}
}

// EnsureNoTokens Ensure provided fields have no tokens.
func (config *Package) EnsureNoTokens(fields []string) error {
	return nil
}

func (config *Package) replaceFieldToken(field string, key string, value string) {
	var target reflect.Value
	switch field {
	case "name":
		target = reflect.ValueOf(config.Name)
	case "description":
		target = reflect.ValueOf(config.Description)
	case "version":
		target = reflect.ValueOf(config.Version)
	case "destination":
		target = reflect.ValueOf(config.Destination)
	case "files":
		target = reflect.ValueOf(config.Files)
	case "exclude":
		target = reflect.ValueOf(config.Exclude)
	case "scripts":
		target = reflect.ValueOf(config.Scripts)
	case "dependencies":
		target = reflect.ValueOf(config.Dependencies)
	case "settings":
		target = reflect.ValueOf(config.Settings)
	case "metadata":
		target = reflect.ValueOf(config.Metadata)
	case "override":
		target = reflect.ValueOf(config.Override)
	}

	replaceTokens(target, key, value)
}

func (config *Package) getSettingConfig(key string) (PackageSetting, bool) {
	for setting, settingConfig := range config.Settings {
		if key == setting {
			return settingConfig, true
		}
	}

	return PackageSetting{}, false
}
