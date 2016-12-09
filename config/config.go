package config

// Config Configuration that zeal will use when running commands.
type Config struct {
	BasePath    string
	DataPath    string
	LogsPath    string
	TempPath    string
	DefaultRepo string
	LogLevel    string
}
