package models

// ConfigItem a config item on zeal config
type ConfigItem struct {
	ID string
	Command string
	Extension string
	Name string
	Source string // can come from code or extension
	Options map[string]interface{}
}