package pack

import (
	"github.com/richardheath/zeal/config"
	"github.com/richardheath/zeal/log"
)

// Pack Package using config file.
func Pack(configPath string, destination string, settings map[string]string) error {
	zealConfig := config.Package{}
	err := zealConfig.Load(configPath)
	if err != nil {
		return err
	}

	log.Debugf("Config path: %s", configPath)
	log.Infof("Packaging: %s", zealConfig.Name)

	if zealConfig.Split {
		return splitPackage(zealConfig)
	}

	return createPackage(zealConfig)
}

func createPackage(zealConfig config.Package) error {
	return nil
}

func splitPackage(zealConfig config.Package) error {
	return nil
}
