package jstra

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Serialize(str interface{}) (string, error) {
	var json strings.Builder

	t := reflect.TypeOf(str)
	v := reflect.ValueOf(str)
	n := t.NumField()

	// TODO: Support Struct Pointers
	if t.Kind() != reflect.Struct {
		err := errors.New("type passed is not of type Struct or Struct Pointer")
		return "", err
	}

	json.WriteString("{")
	// iterate through the fields
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)

		json.WriteString("\"" + jsonFormarter(tt.Name) + "\":")

		switch tt.Type.Kind() {
		case reflect.String:
			json.WriteString("\"" + vv.String() + "\"")
		case reflect.Bool:
			json.WriteString(fmt.Sprintf("%v", vv))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			json.WriteString(fmt.Sprintf("%v", vv))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			json.WriteString(fmt.Sprintf("%v", vv))
		case reflect.Float32, reflect.Float64:
			json.WriteString(fmt.Sprintf("%v", vv))
		case reflect.Slice:
			st := tt.Type.Elem()

			json.WriteString("[")

			for x := 0; x < vv.Len(); x++ {
				switch st.Kind() {
				case reflect.String:
					json.WriteString(fmt.Sprintf("\"%v\"", vv.Index(x)))
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					json.WriteString(fmt.Sprintf("%v", vv.Index(x)))
				case reflect.Float32, reflect.Float64:
					json.WriteString(fmt.Sprintf("%v", vv.Index(x)))
				}
				if x < vv.Len()-1 {
					json.WriteString(",")
				}
			}

			json.WriteString("]")

		}

		if i < n-1 {
			json.WriteString(",")
		}
	}

	json.WriteString("}")

	return json.String(), nil
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
