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
	act, ok := Serialize(Person{Name: "John"})

	exp2 := "{\"name\":\"John\",\"location\":\"Thika\"}"
	act2, ok2 := Serialize(Person2{Name: "John", Location: "Thika"})

	if !ok {
		t.Error("An Error Occured on TestWithOnlyStrings")
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}

	if !ok2 {
		t.Error("An Error Occured on Person2")
	}

	if exp2 != act2 {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act2, exp2)
	}

}

func TestWithStringsandBool(t *testing.T) {
	exp := "{\"typeOfWeather\":\"Sunny\",\"isHumid\":false}"
	act, ok := Serialize(Weather{TypeOfWeather: "Sunny", IsHumid: false})

	if !ok {
		t.Error("An Error Occured on TestWithStringsandBool")
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}
