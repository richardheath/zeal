package data

import (
	"github.com/richardheath/zeal/models"
)

// StateDataHandler Handler for state data
type StateDataHandler interface {
	ListPackages()
	LoadPackage(packageKey string) (PackageDataHandler, error)
}

// PackageDataHandler Interface to load package information
type PackageDataHandler interface {
	GetInfo() (models.PackageInfo, error)

	// Dependency
	GetDependencies() ([]models.PackageDependency, error)
	GetDependency(packageKey string) (models.PackageDependency, error)
	RemoveDependency(packageKey string) error
	AddDependency(models.PackageDependency) error

	// Variables
	GetVariables() ([]models.PackageVariable, error)
	GetVariable(key string) (models.PackageVariable, error)
	SetVariable(models.PackageVariable) error
	ClearCommandVariables(command string) error

	// Config
	AddConfigItem(models.ConfigItem) error
	GetConfig() (models.ConfigItem, error)
	ClearCommandConfigs(models.ConfigItem)

	// 
	GetEventHooks() ([]interface{}, error)
}

type PackageData struct {
	stateDB StateDataHandler
	packageKey string
}
