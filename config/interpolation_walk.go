package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/mitchellh/reflectwalk"
)

const UnknownVariableValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

type interpolationWalker struct {
	key         []string
	lastValue   reflect.Value
	walkerFunc  interpolationWalkerFunc
	Replace     bool
	loc         reflectwalk.Location
	cs          []reflect.Value
	csKey       []reflect.Value
	csData      interface{}
	sliceIndex  []int
	unknownKeys []string
}

func (w *interpolationWalker) Enter(loc reflectwalk.Location) error {
	w.loc = loc
	return nil
}

func (w *interpolationWalker) Exit(loc reflectwalk.Location) error {
	w.loc = reflectwalk.None

	switch loc {
	case reflectwalk.Map:
		w.cs = w.cs[:len(w.cs)-1]
	case reflectwalk.MapValue:
		w.key = w.key[:len(w.key)-1]
		w.csKey = w.csKey[:len(w.csKey)-1]
	case reflectwalk.Slice:
		// Split any values that need to be split
		w.splitSlice()
		w.cs = w.cs[:len(w.cs)-1]
	case reflectwalk.SliceElem:
		w.csKey = w.csKey[:len(w.csKey)-1]
		w.sliceIndex = w.sliceIndex[:len(w.sliceIndex)-1]
	}

	return nil
}

func (w *interpolationWalker) Map(m reflect.Value) error {
	w.cs = append(w.cs, m)
	return nil
}

func (w *interpolationWalker) MapElem(m, k, v reflect.Value) error {
	w.csData = k
	w.csKey = append(w.csKey, k)

	if l := len(w.sliceIndex); l > 0 {
		w.key = append(w.key, fmt.Sprintf("%d.%s", w.sliceIndex[l-1], k.String()))
	} else {
		w.key = append(w.key, k.String())
	}

	w.lastValue = v
	return nil
}

func (w *interpolationWalker) Slice(s reflect.Value) error {
	w.cs = append(w.cs, s)
	return nil
}

func (w *interpolationWalker) SliceElem(i int, elem reflect.Value) error {
	w.csKey = append(w.csKey, reflect.ValueOf(i))
	w.sliceIndex = append(w.sliceIndex, i)
	return nil
}

type interpolationWalkerFunc func(ast.Node) (interface{}, error)

func (w *interpolationWalker) Primitive(v reflect.Value) error {
	setV := v
	fmt.Printf("  %v\n", setV)
	// We only care about strings
	if v.Kind() == reflect.Interface {
		setV = v
		v = v.Elem()
	}
	if v.Kind() != reflect.String {
		return nil
	}

	astRoot, err := hil.Parse(v.String())
	if err != nil {
		return err
	}

	// If the AST we got is just a literal string value with the same
	// value then we ignore it. We have to check if its the same value
	// because it is possible to input a string, get out a string, and
	// have it be different. For example: "foo-$${bar}" turns into
	// "foo-${bar}"
	if n, ok := astRoot.(*ast.LiteralNode); ok {
		if s, ok := n.Value.(string); ok && s == v.String() {
			return nil
		}
	}

	if w.walkerFunc == nil {
		return nil
	}

	replaceVal, err := w.walkerFunc(astRoot)
	if err != nil {
		return fmt.Errorf(
			"%s in:\n\n%s",
			err, v.String())
	}

	if !w.Replace {
		return nil
	}

	// We need to determine if we need to remove this element
	// if the result contains any "UnknownVariableValue" which is
	// set if it is computed. This behavior is different if we're
	// splitting (in a SliceElem) or not.
	remove := false
	if w.loc == reflectwalk.SliceElem {
		switch typedReplaceVal := replaceVal.(type) {
		case string:
			if typedReplaceVal == UnknownVariableValue {
				remove = true
			}
		case []interface{}:
			if hasUnknownValue(typedReplaceVal) {
				remove = true
			}
		}
	} else if replaceVal == UnknownVariableValue {
		remove = true
	}

	if remove {
		w.unknownKeys = append(w.unknownKeys, strings.Join(w.key, "."))
	}

	resultVal := reflect.ValueOf(replaceVal)
	switch w.loc {
	case reflectwalk.MapKey:
		m := w.cs[len(w.cs)-1]

		// Delete the old value
		var zero reflect.Value
		m.SetMapIndex(w.csData.(reflect.Value), zero)

		// Set the new key with the existing value
		m.SetMapIndex(resultVal, w.lastValue)

		// Set the key to be the new key
		w.csData = resultVal
	case reflectwalk.MapValue:
		// If we're in a map, then the only way to set a map value is
		// to set it directly.
		m := w.cs[len(w.cs)-1]
		mk := w.csData.(reflect.Value)
		m.SetMapIndex(mk, resultVal)
	default:
		// Otherwise, we should be addressable
		fmt.Printf("resultVal: %v", resultVal)
		setV.Set(resultVal)
	}

	return nil
}

func hasUnknownValue(variable []interface{}) bool {
	for _, value := range variable {
		if strVal, ok := value.(string); ok {
			if strVal == UnknownVariableValue {
				return true
			}
		}
	}
	return false
}

func (w *interpolationWalker) splitSlice() {
	raw := w.cs[len(w.cs)-1]

	var s []interface{}
	switch v := raw.Interface().(type) {
	case []interface{}:
		s = v
	case []map[string]interface{}:
		return
	}

	split := false
	for _, val := range s {
		if varVal, ok := val.(ast.Variable); ok && varVal.Type == ast.TypeList {
			split = true
		}
		if _, ok := val.([]interface{}); ok {
			split = true
		}
	}

	if !split {
		return
	}

	result := make([]interface{}, 0)
	for _, v := range s {
		switch val := v.(type) {
		case ast.Variable:
			switch val.Type {
			case ast.TypeList:
				elements := val.Value.([]ast.Variable)
				for _, element := range elements {
					result = append(result, element.Value)
				}
			default:
				result = append(result, val.Value)
			}
		case []interface{}:
			for _, element := range val {
				result = append(result, element)
			}
		default:
			result = append(result, v)
		}
	}

	w.replaceCurrent(reflect.ValueOf(result))
}

func (w *interpolationWalker) replaceCurrent(v reflect.Value) {
	// if we don't have at least 2 values, we're not going to find a map, but
	// we could panic.
	if len(w.cs) < 2 {
		return
	}

	c := w.cs[len(w.cs)-2]
	switch c.Kind() {
	case reflect.Map:
		// Get the key and delete it
		k := w.csKey[len(w.csKey)-1]
		c.SetMapIndex(k, v)
	}
}
