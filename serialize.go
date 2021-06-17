package jstra

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Serialize(str interface{}) (string, error) {
	var json string

	t := reflect.TypeOf(str)
	v := reflect.ValueOf(str)
	n := t.NumField()

	// TODO: Support Struct Pointers
	if t.Kind() != reflect.Struct {
		err := errors.New("type passed is not of type Struct or Struct Pointer")
		return "", err
	}

	json += "{"
	// iterate through the fields
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)

		json += "\"" + jsonFormarter(tt.Name) + "\":"

		switch tt.Type.Kind() {
		case reflect.String:
			json += "\"" + vv.String() + "\""
		case reflect.Bool:
			json += fmt.Sprintf("%v", vv)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			json += fmt.Sprintf("%v", vv)
		}

		if i < n-1 {
			json += ","
		}
	}

	json += "}"

	return json, nil
}

func jsonFormarter(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	b := []byte(s)
	r := b[1:]
	lc := bytes.ToLower([]byte{b[0]})
	return string(bytes.Join([][]byte{lc, r}, nil))
}
