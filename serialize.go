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
		js.arrays2Json(v.Interface())
	} else if t.Kind() == reflect.Ptr {
		p := reflect.Indirect(v)
		js.serializer(p.Interface())

	} else if t.Kind() == reflect.Struct {
		n := t.NumField()

		js.json.WriteByte('{')
		// iterate through the fields
		for i := 0; i < n; i++ {
			tt := t.Field(i)
			vv := v.Field(i)

			js.json.WriteByte('"')
			js.json.Write(jsonFormarter(tt.Name))
			js.json.WriteByte('"')
			js.json.WriteByte(':')

			switch tt.Type.Kind() {

			case reflect.String:
				js.json.WriteByte('"')
				s := []byte(vv.String())
				js.json.Write(s)
				js.json.WriteByte('"')

			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				js.json.WriteString(fmt.Sprintf("%v", vv))
			case reflect.Slice:
				js.arrays2Json(vv.Interface())

			case reflect.Struct:
				js.serializer(vv.Interface())
			}

			if i < n-1 {
				js.json.WriteByte(',')

			}
		}

		js.json.WriteByte('}')
	} else {
		err := errors.New("type passed is not of type Struct or Struct Pointer")
		return "", err
	}

	return js.json.String(), nil
}

func (js *jstraSerialize) arrays2Json(str interface{}) {

	t := reflect.TypeOf(str)
	v := reflect.ValueOf(str)

	k := t.Elem().Kind()

	js.json.WriteByte('[')
	defer js.json.WriteByte(']')
	for x := 0; x < v.Len(); x++ {
		switch k {
		case reflect.String:
			// js.json.WriteString("\"" + v.Index(x).String() + "\"")
			js.json.WriteByte('"')
			s := []byte(v.Index(x).String())
			js.json.Write(s)
			js.json.WriteByte('"')
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			js.json.WriteString(fmt.Sprintf("%v", v.Index(x)))
		case reflect.Struct:
			js.serializer(v.Index(x).Interface())
		}

		if x < v.Len()-1 {
			js.json.WriteByte(',')
		}
	}

}

func jsonFormarter(s string) []byte {
	b := []byte(s)

	if len(s) < 2 {
		return bytes.ToLower(b)
	}

	r := b[1:]
	lc := bytes.ToLower([]byte{b[0]})
	return bytes.Join([][]byte{lc, r}, nil)
}
