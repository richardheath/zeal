package config

// Package Zeal package configuration.
type Package struct {
	name         string
	description  string
	version      string
	keep         string
	split        bool
	author       string
	destination  string
	files        map[string]string
	exclude      []string
	scripts      map[string][]string
	dependencies string
	settings     []PackageSetting
	metadata     map[string]string
	override     map[string]Package
}

// PackageSetting Zeal package setting configuration.
type PackageSetting struct {
	defaultValue string
	description  string
}
