package config

import (
	"strings"
	"fmt"

	"github.com/hashicorp/hil/ast"
	"github.com/mitchellh/reflectwalk"
	"github.com/richardheath/zeal/data"
)

// Config type
type Config struct {
	items []ConfigItem
}

// ConfigItem a config item on zeal config
type ConfigItem struct {
	ID        string
	Command   string
	Extension string
	Name      string
	Source    string // can come from code or extension
	Options   map[string]interface{}
}

func (config *Config) SetItems(items []ConfigItem) {
	config.items = items
}

// Evaluate Evaluate config items and reorder based on interpolation.
func (config *Config) Evaluate(packageData data.PackageDataHandler) error {
	dependencies := map[string]string{}
	fmt.Println(dependencies)

	walkerFunc := func(node ast.Node) (interface{}, error) {
		fmt.Println("EVAL!!!")
		variableUsed := getVariablesUsed(fmt.Sprintf("%v", node))
		fmt.Println(variableUsed)
		return "", nil
	}

	walker := &interpolationWalker{
		walkerFunc: walkerFunc,
		Replace:    false,
	}

	err := reflectwalk.Walk(config.items, walker)
	if err != nil {
		return err
	}

	return nil
}

// Filter Filter config items. Use empty string to skip filter on field
// Sample config.Filter("build", "", "") // Only filter builds
func (config *Config) Filter(command string, extension string, name string) []ConfigItem {
	result := []ConfigItem{}
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

func getConfigItemUsed(raw string) []string {
	varPrefix := "Variable("
	varPrefixLen := len(varPrefix)

	items := []string
	pos := 0
	len := len(raw)
	varIndex := strings.Index(raw, varPrefix)
	for varIndex > -1 {
		raw = raw[varIndex+varPrefixLen:]
		endIndex := strings.Index(raw, ")")
		items = append(items, )


		varIndex := strings.Index(raw, varPrefix)
	}
	return nil
}
