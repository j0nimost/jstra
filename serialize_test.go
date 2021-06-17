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

type Catalogue struct {
	Id        int64
	Item      string
	IsInStock bool
	Quantity  int32
	Review    int8
}

type Catalogue2 struct {
	Id        uint64
	Item      string
	IsInStock bool
	Quantity  int32
	Review    uint8
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

func TestWithIntegers(t *testing.T) {
	exp := "{\"id\":-8223372036854775807,\"item\":\"Yves Saint Laurent\",\"isInStock\":true,\"quantity\":45,\"review\":-5}"
	exp2 := "{\"id\":8223372036854775807,\"item\":\"Yves Saint Laurent\",\"isInStock\":true,\"quantity\":45,\"review\":5}"

	act, err := Serialize(Catalogue{Id: -8223372036854775807, Item: "Yves Saint Laurent", IsInStock: true, Quantity: 45, Review: -5})
	act2, err2 := Serialize(Catalogue{Id: 8223372036854775807, Item: "Yves Saint Laurent", IsInStock: true, Quantity: 45, Review: 5})

	if err != nil {
		t.Errorf(err.Error())
	}

	if err2 != nil {
		t.Errorf(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}

	if exp2 != act2 {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act2, exp2)
	}
}
