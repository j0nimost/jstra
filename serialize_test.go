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
	Books      []string
	TaxNumbers []uint32
	Prices     []int32
	Distances  []float32
}

type B struct {
	X int
	Y int
}

type A struct {
	Name string
	B
}

type C struct {
	IsScalar bool
	A
	IsVector bool // to test if order persists
}

type D struct {
	Bs []B
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
	exp := "{\"books\":[\"Old McDonald Had A Farm\",\"Calculus I\"],\"taxNumbers\":[123455,235555,6643233,664433],\"prices\":[-231,334,64645,23232,55,4,232,54,232,544,3232,343],\"distances\":[12.5,454.56,343.2546,45434.562]}"

	book := []string{"Old McDonald Had A Farm", "Calculus I"}
	taxNumbers := []uint32{123455, 235555, 6643233, 664433}
	prices := []int32{-231, 334, 64645, 23232, 55, 4, 232, 54, 232, 544, 3232, 343}
	distances := []float32{12.5, 454.56, 343.2546, 45434.562}

	act, err := Serialize(SliceType{Books: book, TaxNumbers: taxNumbers, Prices: prices, Distances: distances})

	if err != nil {
		t.Error(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct: Resulted to %s instead of %s\n", act, exp)
	}
}

func TestWithNestedStructs(t *testing.T) {
	exp := "{\"name\":\"Triange\",\"b\":{\"x\":12,\"y\":24}}"
	exp1 := "{\"isScalar\":true,\"a\":{\"name\":\"Triange\",\"b\":{\"x\":12,\"y\":24}},\"isVector\":false}"
	exp2 := "{\"bs\":[{\"x\":12,\"y\":24},{\"x\":13,\"y\":14},{\"x\":15,\"y\":64}]}"
	b := B{X: 12, Y: 24}
	a := A{Name: "Triange", B: b}

	c := C{IsScalar: true, A: a, IsVector: false}

	d := D{Bs: []B{
		{X: 12, Y: 24},
		{X: 13, Y: 14},
		{X: 15, Y: 64},
	}}

	act, err := Serialize(a)
	act1, err1 := Serialize(c)
	act2, err2 := Serialize(d)

	if err != nil {
		t.Error(err.Error())
	}

	if err1 != nil {
		t.Error(err1.Error())
	}

	if err2 != nil {
		t.Error(err2.Error())
	}

	if exp != act {
		t.Errorf("Serializing Nested Structs: Resulted to %s instead of %s\n", act, exp)
	}

	if exp1 != act1 {
		t.Errorf("Serializing Nested Structs: Resulted to %s instead of %s\n", act1, exp1)
	}

	if exp2 != act2 {
		t.Errorf("Serializing Nested Structs: Resulted to %s instead of %s\n", act2, exp2)
	}
}

func TestWithStructPointer(t *testing.T) {
	exp := "{\"x\":23,\"y\":43}"
	b := B{X: 23, Y: 43}

	act, err := Serialize(&b)

	if err != nil {
		t.Error(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Nested Structs: Resulted to %s instead of %s\n", act, exp)
	}
}

func TestWithStructSlices(t *testing.T) {
	exp := "[{\"x\":23,\"y\":43},{\"x\":12,\"y\":24},{\"x\":13,\"y\":14},{\"x\":15,\"y\":64}]"
	exp1 := "[{\"x\":23,\"y\":43},{\"x\":12,\"y\":24},{\"x\":13,\"y\":14},{\"x\":15,\"y\":64}]"
	b := []B{
		{X: 23, Y: 43},
		{X: 12, Y: 24},
		{X: 13, Y: 14},
		{X: 15, Y: 64}}

	act, err := Serialize(b)
	act1, err1 := Serialize(&b)

	if err != nil {
		t.Error(err.Error())
	}

	if err1 != nil {
		t.Error(err.Error())
	}

	if exp != act {
		t.Errorf("Serializing Struct Slices: Resulted to %s instead of %s\n", act, exp)
	}

	if exp1 != act1 {
		t.Errorf("Serializing Struct Slices: Resulted to %s instead of %s\n", act, exp)
	}
}
