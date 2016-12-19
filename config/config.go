package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

// Config Configuration that zeal will use when running commands.
type Config struct {
	DataPath    string
	LogsPath    string
	TempPath    string
	DefaultRepo string
	MaxLogsSize int
	MaxTempSize int
}

const (
	configFileName = "zeal.json"
	zealFolder     = "zeal"
	dataFolder     = "data"
	logsFolder     = "logs"
	tempFolder     = "temp"
	defaultRepo    = "files"
)

var config *Config

// Initialize Initialize zeal configuration.
func Initialize() error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = loadConfigFile(configFile)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig Get zeal global configuration.
func GetConfig() *Config {
	return config
}

func loadConfigFile(filePath string) error {
	defaultConfig := getDefaultConfig()
	config = &defaultConfig

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

func getDefaultConfig() Config {
	usr, _ := user.Current()
	base := path.Join(usr.HomeDir, zealFolder)
	return Config{
		DataPath:    path.Join(base, dataFolder),
		LogsPath:    path.Join(base, logsFolder),
		TempPath:    path.Join(base, tempFolder),
		DefaultRepo: defaultRepo,
		MaxLogsSize: 1000,
		MaxTempSize: 10000,
	}
}

func getConfigFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(usr.HomeDir, configFileName), nil
}
