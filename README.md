## jstra
A small Library Which Parses Structs to Json

[![Build Status](https://travis-ci.com/j0nimost/jstra.svg?token=zBU3HpXnQ9WSEWzAzXky&branch=main)](https://travis-ci.com/j0nimost/jstra)

### Types Supported
- Bool
- Ints
- UInts
- Strings
- Slices (ints, uints, strings, floats, structs)
- Nested Structs
- Struct Pointers
- Struct Slices

### Sample
Given a Struct

```go

type Person struct {
	Name     string
	Age      uint
	Contacts []string
	NetPay   float64
}
```

You implement the package like so

```go

func main() {

	fmt.Println("Serialize")
	p := Person{Name: "Ken Alex", Age: 24, Contacts: []string{"02323232", "23232533"}, NetPay: 2000.50}
	json, err := jstra.Serialize(&p)

	if err != nil {
		panic(err)
	}

	fmt.Println(json)
}
```


It gives the following output

```json
{
	"name": "Ken Alex",
	"age": 24,
	"contacts": ["02323232", "23232533"],
	"netPay": 2000.5
}
```

### Contribution
FORK and HACK 

### Author
John Nyingi
