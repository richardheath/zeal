package pack

import (
	"github.com/richardheath/zeal/config"
	"github.com/richardheath/zeal/log"
)

// Pack Package using config file.
func Pack(configPath string, destination string, settings map[string]string) error {
	log.Debugf("Config path: %s", configPath)
	log.Debugf("Settings: %v", settings)

	zealConfig, err := processConfig(configPath, settings)
	if err != nil {
		return err
	}

	log.Infof("Packaging: %s", zealConfig.Name)
	if zealConfig.Split {
		return splitPackage(zealConfig)
	}

	return createPackage(zealConfig)
}

func processConfig(configPath string, settings map[string]string) (config.Package, error) {
	zealConfig := config.Package{}
	err := zealConfig.Load(configPath)
	if err != nil {
		return zealConfig, err
	}

	requiredFields := []string{
		"name",
		"description",
		"version",
		"author",
		"dependencies",
		"metadata",
		"settings",
	}

	zealConfig.ReplaceTokens(settings)
	zealConfig.ReplaceWithDefaultTokens(requiredFields)
	err = zealConfig.EnsureNoTokens(requiredFields)
	return zealConfig, err
}

func createPackage(zealConfig config.Package) error {
	return nil
}

func splitPackage(zealConfig config.Package) error {
	return nil
}
