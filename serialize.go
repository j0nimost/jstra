package jstra

import (
	"reflect"
	"strings"
)

func Serialize(str interface{}) (string, bool) {
	var json string

	t := reflect.TypeOf(str)
	v := reflect.ValueOf(str)

	if t.Kind() != reflect.Struct {
		return "", false
	}

	json += "{"
	// iterate through the fields
	for i := 0; i < t.NumField(); i++ {
		tt := t.Field(i)
		vv := v.Field(i)

		json += "\"" + strings.ToLower(tt.Name) + "\":"

		switch tt.Type.Kind() {
		case reflect.String:
			json += "\"" + vv.String() + "\""
		case reflect.Bool:
			json += vv.String()
		}
		// get tt value, type and name
		// concatenate into a string
	}

	json += "}"

	return json, true
}
