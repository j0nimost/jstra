package jstra

import "testing"

type Person struct {
	Name string
}

type Person2 struct {
	Name     string
	Location string
}

type Weather struct {
	TypeOfWeather string
	IsHumid       bool
}

func TestWithOnlyStrings(t *testing.T) {

	exp := "{\"name\":\"John\"}"
	act, err := Serialize(Person{Name: "John"})

	exp2 := "{\"name\":\"John\",\"location\":\"Thika\"}"
	act2, err2 := Serialize(Person2{Name: "John", Location: "Thika"})

	if err != nil {
		t.Errorf(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}

	if err2 != nil {
		t.Error("An Error Occured on Person2")
	}

	if exp2 != act2 {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act2, exp2)
	}

}

func TestWithStringsandBool(t *testing.T) {
	exp := "{\"typeOfWeather\":\"Sunny\",\"isHumid\":false}"
	act, err := Serialize(Weather{TypeOfWeather: "Sunny", IsHumid: false})

	if err != nil {
		t.Errorf(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}
