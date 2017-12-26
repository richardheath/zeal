package models


// PackageInfo Settings information
type PackageInfo struct {
	Name string
    Description string
    Version string
    Author string
    UpdateRule string
}

// PackageDependency Dependency information
type PackageDependency struct {
	PackageKey string
	Name string
	Version string
}

type PackageVariable struct {
	Key string
	Value interface{}
}