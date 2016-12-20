package config

type Package struct {
	name             string
	description      string
	version          string
	multiVersionRule string
	splitPerOS       bool
	author           string
	contents         PackageContents
	scripts          map[string][]string
	dependencies     string
	settings         []PackageSetting
	metadata         map[string]string
	osOverride       map[string]Package
}

type PackageContents struct {
	basePath    string
	destination string
	files       string
	exclude     string
}

type PackageSetting struct {
	defaultValue string
	description  string
}
