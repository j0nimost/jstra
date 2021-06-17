package jstra

import "testing"

type StringType struct {
	Name     string
	Location string
}

type BoolType struct {
	IsHumid bool
}

type IntType struct {
	Id       int64
	Quantity int32
	Review   int8
}

type UIntType struct {
	Id       uint64
	Quantity uint32
	Review   uint8
}

type FloatType struct {
	Credit float32
	Debit  float32
	Pi     float64
}

type SliceType struct {
	Books []string
}

func TestWithOnlyStrings(t *testing.T) {

	exp := "{\"name\":\"John\",\"location\":\"\"}"
	act, err := Serialize(StringType{Name: "John"})

	exp2 := "{\"name\":\"John\",\"location\":\"Thika\"}"
	act2, err2 := Serialize(StringType{Name: "John", Location: "Thika"})

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
	exp := "{\"isHumid\":false}"
	act, err := Serialize(BoolType{IsHumid: false})

	if err != nil {
		t.Errorf(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}

func TestWithIntegers(t *testing.T) {
	exp := "{\"id\":-8223372036854775807,\"quantity\":45,\"review\":-5}"
	exp2 := "{\"id\":8223372036854775807,\"quantity\":45,\"review\":5}"

	act, err := Serialize(IntType{Id: -8223372036854775807, Quantity: 45, Review: -5})
	act2, err2 := Serialize(UIntType{Id: 8223372036854775807, Quantity: 45, Review: 5})

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

func TestWithFloats(t *testing.T) {
	exp := "{\"credit\":1538.65,\"debit\":545.35,\"pi\":3.142857142857143}"

	p := 22.0 / 7.0
	act, err := Serialize(FloatType{Credit: 1538.65, Debit: 545.35, Pi: p})

	if err != nil {
		t.Errorf(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}

func TestWithSlices(t *testing.T) {
	exp := "{\"books\":[\"Old McDonald Had A Farm\",\"Calculus I\"]}"

	book := []string{"Old McDonald Had A Farm", "Calculus I"}
	act, err := Serialize(SliceType{Books: book})

	if err != nil {
		t.Error(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}
