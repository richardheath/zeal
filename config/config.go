package config

// Config type
type Config struct {
	items []Item
}

// Item a config item on zeal config
type Item struct {
	ID string
	Command string
	Extension string
	Name string
	Options map[string]interface{}
}

// Filter Filter config items. Use empty string to skip filter on field
// Sample config.Filter("build", "", "") // Only filter builds
func (config *Config) Filter(command string, extension string, name string) []Item {
	result := []Item{}
	if command != "" {
		for _, item := range config.items {
			if command == item.Command {
				result = append(result, item)
			}
		}
	}

	if extension != "" {
		for _, item := range config.items {
			if extension == item.Extension {
				result = append(result, item)
			}
		}
	}

	if name != "" {
		for _, item := range config.items {
			if name == item.Name {
				result = append(result, item)
			}
		}
	}

	return result
}
