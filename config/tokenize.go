package config

import (
	"reflect"
	"strings"
)

func replaceTokens(target reflect.Value, token string, value string) {
	switch target.Kind() {
	case reflect.String:
		replaceStringWithValue(target, token, value)
	case reflect.Map:
		replaceTokensInMap(target, token, value)
	case reflect.Struct:
		replaceTokensInStruct(target, token, value)
	case reflect.Array:
		replaceTokensInCollection(target, token, value)
	case reflect.Slice:
		replaceTokensInCollection(target, token, value)
	}
}

func replaceTokensInCollection(target reflect.Value, token string, value string) {
	for i := 0; i < target.Len(); i++ {
		replaceTokens(target.Index(i), token, value)
	}
}

func replaceTokensInMap(target reflect.Value, token string, value string) {
	for _, mapKey := range target.MapKeys() {
		mapValue := target.MapIndex(mapKey)
		if mapValue.Kind() == reflect.String {
			result := replaceToken(mapValue.String(), token, value)
			target.SetMapIndex(mapKey, reflect.ValueOf(result))
		} else {
			replaceTokens(target.MapIndex(mapKey), token, value)
		}
	}
}

func replaceTokensInStruct(target reflect.Value, token string, value string) {
	for i := 0; i < target.NumField(); i++ {
		replaceTokens(target.Field(i), token, value)
	}
}

func replaceStringWithValue(target reflect.Value, key string, value string) {
	if !target.CanSet() {
		return
	}

	var result = strings.Replace(target.String(), key, value, -1)
	target.SetString(result)
}

func replaceToken(raw string, key string, value string) string {
	return strings.Replace(raw, key, value, -1)
}
