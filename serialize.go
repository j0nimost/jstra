package jstra

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type jstraSerialize struct {
	json strings.Builder
}

func Serialize(str interface{}) (string, error) {

	js := newJstraSerialize()

	return js.serializer(str)

}

func newJstraSerialize() *jstraSerialize {
	return &jstraSerialize{json: strings.Builder{}}
}

func (js *jstraSerialize) serializer(str interface{}) (string, error) {

	t := reflect.TypeOf(str)
	v := reflect.ValueOf(str)

	if t.Kind() == reflect.Slice {
		st := t.Elem()

		js.json.WriteString("[")
		for x := 0; x < v.Len(); x++ {
			switch st.Kind() {
			case reflect.Struct:
				js.serializer(v.Index(x).Interface())
			default:
				err := errors.New("slice type passed is not of type Struct or Struct Pointer")
				return "", err
			}
			if x < v.Len()-1 {
				js.json.WriteString(",")
			}

		}
		js.json.WriteString("]")

	} else if t.Kind() == reflect.Ptr {
		p := reflect.Indirect(v)
		js.serializer(p.Interface())

	} else if t.Kind() == reflect.Struct {
		n := t.NumField()

		js.json.WriteString("{")
		// iterate through the fields
		for i := 0; i < n; i++ {
			tt := t.Field(i)
			vv := v.Field(i)

			js.json.WriteString("\"" + jsonFormarter(tt.Name) + "\":")

			switch tt.Type.Kind() {
			case reflect.String:
				js.json.WriteString("\"" + vv.String() + "\"")
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				js.json.WriteString(fmt.Sprintf("%v", vv))
			case reflect.Slice:
				st := tt.Type.Elem()

				js.json.WriteString("[")

				for x := 0; x < vv.Len(); x++ {
					switch st.Kind() {
					case reflect.String:
						js.json.WriteString(fmt.Sprintf("\"%v\"", vv.Index(x)))
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
						reflect.Float32, reflect.Float64:
						js.json.WriteString(fmt.Sprintf("%v", vv.Index(x)))
					case reflect.Struct:
						js.serializer(vv.Index(x).Interface())
					}
					if x < vv.Len()-1 {
						js.json.WriteString(",")
					}

				}

				js.json.WriteString("]")

			case reflect.Struct:
				js.serializer(vv.Interface())
			}

			if i < n-1 {
				js.json.WriteString(",")

			}
		}

		js.json.WriteString("}")
	} else {
		err := errors.New("type passed is not of type Struct or Struct Pointer")
		return "", err
	}

	return js.json.String(), nil
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
