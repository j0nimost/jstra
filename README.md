## jstra
Simple Package Prof of Concept Library which Serializes Structs to Json

[![Build Status](https://travis-ci.com/j0nimost/jstra.svg?token=zBU3HpXnQ9WSEWzAzXky&branch=main)](https://travis-ci.com/j0nimost/jstra)
### Sample
Given a Struct

```go

type Person struct {
	Name string
}
```

You implement the package like so

```go

func main() {
	s, k := jstra.Serialize(Person{Name: "John"})

	if k != nil {
		fmt.Println(k)
	}

	fmt.Println(s)
}

```


It gives the following output

```json

{"name":"John"}
```

### Types Supported
- Bool
- Ints
- UInts
- Strings
- Slices (ints, uints, floats, structs)
- Nested Structs

### Contribution
FORK and HACK 

### Author
John Nyingi
