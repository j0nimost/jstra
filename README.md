## jstra
Simple Package which Serializes Struct to Json

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

	if !k {
		fmt.Print("Error\n")
		return
	}

	fmt.Println(s)
}

```


It gives the following output

```json

{"name":"John"}
```

### Contribution
FORK and HACK

### Author
John Nyingi
