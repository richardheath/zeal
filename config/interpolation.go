package config

import (
	"fmt"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/mitchellh/reflectwalk"
	"github.com/richardheath/zeal/data"
)

// Interpolate Run interpolation using package data handler.
func (item *ConfigItem) Interpolate(packageData data.PackageDataHandler) error {
	//vars := map[string]ast.Variable{}
	//config := langEvalConfig(vars)
	config := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap: map[string]ast.Variable{
				"build.go.output": ast.Variable{
					Type:  ast.TypeString,
					Value: "interpolate",
				},
				"build.go.test": ast.Variable{
					Type:  ast.TypeString,
					Value: "test",
				},
			},
			FuncMap: map[string]ast.Function{},
		},
	}

	walkerFunc := func(node ast.Node) (interface{}, error) {
		fmt.Println(node)
		fmt.Println(node.Pos())
		result, err := hil.Eval(node, config)
		if err != nil {
			return "", err
		}
		fmt.Println(result)
		return result.Value, nil
	}

	walker := &interpolationWalker{walkerFunc: walkerFunc}
	err := reflectwalk.Walk(item.Options, walker)
	if err != nil {
		return err
	}

	return nil
}
