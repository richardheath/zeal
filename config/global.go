package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

// ZealConfig Configuration that zeal will use when running commands.
type ZealConfig struct {
	DataPath    string `json:"dataPath"`
	LogsPath    string `json:"logsPath"`
	TempPath    string `json:"tempPath"`
	DefaultRepo string `json:"defaultRepo"`
	MaxLogsSize int    `json:"maxLogsSize"`
	MaxTempSize int    `json:"maxTempSize"`
}

const (
	configFileName = "zeal.json"
	zealFolder     = "zeal"
	dataFolder     = "data"
	logsFolder     = "logs"
	tempFolder     = "temp"
	defaultRepo    = "files"
)

var zealConfig *ZealConfig

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

// GetZealConfig Get zeal global configuration.
func GetZealConfig() *ZealConfig {
	return zealConfig
}

func loadConfigFile(filePath string) error {
	defaultConfig := getDefaultConfig()
	zealConfig = &defaultConfig

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	rawConfig, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	json.Unmarshal(rawConfig, zealConfig)
	return nil
}

func getDefaultConfig() ZealConfig {
	usr, _ := user.Current()
	base := path.Join(usr.HomeDir, zealFolder)
	return ZealConfig{
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
