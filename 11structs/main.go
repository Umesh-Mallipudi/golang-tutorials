package main

import (
	"fmt"
)

// 1. Define a struct
type Person struct {
	name string
	age  int
}

func main() {
	// Way 1: Struct literal with field names (Best Practice)
	p1 := Person{name: "Alice", age: 30}
	fmt.Println("1. p1:", p1.name, p1.age)

	// Way 2: Struct literal without field names (Not Recommended)
	p2 := Person{"Bob", 25}
	fmt.Println("2. p2:", p2.name, p2.age)

	// Way 3: Declare and assign fields one by one
	var p3 Person
	p3.name = "Charlie"
	p3.age = 40
	fmt.Println("3. p3:", p3.name, p3.age)

	// Way 4: Pointer to struct using & (shortcut)
	p4 := &Person{name: "David", age: 22}
	fmt.Println("4. p4:", p4.name, p4.age)

	// Way 5: Pointer to struct using new()
	p5 := new(Person) // p5 is of type *Person
	p5.name = "Eve"
	p5.age = 28
	fmt.Println("5. p5:", p5.name, p5.age)

	// Way 6: Anonymous struct
	phone := struct {
		model string
		price int
	}{
		model: "XPhone",
		price: 999,
	}
	fmt.Println("6. Anonymous struct:", phone.model, phone.price)
}
