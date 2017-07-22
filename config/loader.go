package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl"
)

const zealExtension = ".tf"

// LoadDir Load zeal config files directory
func LoadDir(dirPath string) (*Config, error) {
	config := Config{}
	files, err := getConfigFiles(dirPath)
	if err != nil {
		return &config, err
	}

	data, _ := mergeTextFiles(dirPath, files)
	if err != nil {
		return &config, err
	}

	return generateConfig(data)
}

func getConfigFiles(dirPath string) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	dirStat, err := dir.Stat()
	if err != nil {
		return nil, err
	}
	if !dirStat.IsDir() {
		return nil, fmt.Errorf(
			"configuration path must be a directory: %s",
			dirPath)
	}

	var files []string
	filepath.Walk(dirPath, func(path string, fi os.FileInfo, _ error) error {
		if !fi.IsDir() {
			match, err := regexp.MatchString(zealExtension, fi.Name())
			if err == nil && match {
				files = append(files, fi.Name())
			}
		}
		return nil
	})

	return files, nil
}

func generateConfig(data string) (*Config, error) {
	config := Config{
		items: []Item{},
	}

	hclRoot, err := hcl.Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("Error parsing %s", err)
	}
	
	list, ok := hclRoot.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file doesn't contain a root object")
	}
	
	for _, item := range list.Children().Items {
		configItem := Item{}

		keysLen := len(item.Keys)
		if keysLen > 0 {
			configItem.Command = item.Keys[0].Token.Value().(string)
		}
		if keysLen > 1 {
			configItem.Extension = item.Keys[1].Token.Value().(string)
		}
		if keysLen > 2 {
			configItem.Name = item.Keys[2].Token.Value().(string)
		}

		if _, ok := item.Val.(*ast.ObjectType); !ok {
			return &config, fmt.Errorf(
				"config %s.%s.%s: should be an object",
				configItem.Command,
				configItem.Extension,
				configItem.Name)
		}

		if err := hcl.DecodeObject(&configItem.Options, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading config for %s.%s.%s: %s",
				configItem.Command,
				configItem.Extension,
				configItem.Name,
				err)
		}

		config.items = append(config.items, configItem)
	}

	return &config, nil
}

func mergeTextFiles(basePath string, files []string) (string, error) {
	result := []byte{}
	for _, file := range files {
		content, err := ioutil.ReadFile(basePath + "/" + file)
		if err != nil {
			return "", fmt.Errorf("Error reading %s: %s", file, err)
		}

		result = append(result, content...)
	}
	
	return string(result), nil
}
