Great! Let's dive into **methods** in Go and how they differ from **functions**, followed by all relevant **scenarios of using methods with structs**.

---

## 🔹 Difference Between Functions and Methods in Go

| Feature    | **Function**             | **Method**                                |
| ---------- | ------------------------ | ----------------------------------------- |
| Definition | Declared independently   | Associated with a type (usually a struct) |
| Syntax     | `func functionName(...)` | `func (receiver Type) methodName(...)`    |
| Called On  | Called directly          | Called on an instance of a type           |
| Purpose    | General-purpose logic    | Behavior specific to a type (struct)      |

---

## 🔹 Method Syntax in Go

```go
type Person struct {
    Name string
    Age  int
}

// Method with value receiver
func (p Person) SayHello() {
    fmt.Println("Hello, my name is", p.Name)
}

// Method with pointer receiver
func (p *Person) HaveBirthday() {
    p.Age++
}
```

---

## 🔹 Scenarios of Methods in Structs

Let’s go through the **key scenarios** where methods are used with structs.

---

### ✅ 1. **Method with Value Receiver**

Does **not modify** the original struct. It gets a copy.

```go
func (p Person) Greet() {
    fmt.Println("Hi,", p.Name)
}
```

Usage:

```go
p := Person{"Alice", 25}
p.Greet() // Outputs: Hi, Alice
```

---

### ✅ 2. **Method with Pointer Receiver**

Can **modify** the original struct.

```go
func (p *Person) Birthday() {
    p.Age += 1
}
```

Usage:

```go
p := Person{"Bob", 30}
p.Birthday()
fmt.Println(p.Age) // 31
```

Go automatically handles pointer/value receiver conversions when calling methods.

---

### ✅ 3. **Methods on Custom Types (not just structs)**

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*1.8 + 32
}
```

Usage:

```go
temp := Celsius(30)
fmt.Println(temp.ToFahrenheit()) // 86
```

---

### ✅ 4. **Using Interfaces with Methods**

If a struct implements all methods of an interface, it satisfies the interface.

```go
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}
```

Usage:

```go
var s Speaker = Dog{}
s.Speak() // Woof!
```

---

### ✅ 5. **Embedding Structs and Inheriting Methods**

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
    Animal
}

func (d Dog) Bark() {
    fmt.Println(d.Name, "barks")
}
```

Usage:

```go
d := Dog{Animal{"Rex"}}
d.Speak() // Rex makes a sound
d.Bark()  // Rex barks
```

---

### ✅ 6. **Method Overriding / Shadowing**

If a struct embeds another and defines a method with the same name, the outer method overrides the embedded one.

```go
func (d Dog) Speak() {
    fmt.Println(d.Name, "says Woof!")
}
```

Now `d.Speak()` uses the Dog version, not Animal's.

---

### ✅ 7. **Receiver Type Rules**

* Receiver type must be defined in the **same package**.
* Cannot define a method on a built-in type (like `int` or `string`) directly.

---

## 🧪 Example Bringing All Together

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Info() {
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func (p *Person) GrowUp() {
    p.Age++
}

type Employee struct {
    Person
    Position string
}

func (e Employee) Work() {
    fmt.Println(e.Name, "is working as", e.Position)
}

func main() {
    e := Employee{
        Person:   Person{Name: "Alice", Age: 25},
        Position: "Engineer",
    }

    e.Info()      // Method from embedded struct
    e.Work()      // Method from Employee
    e.GrowUp()    // Pointer receiver method
    e.Info()      // Age is now incremented
}
```

---

## ✅ Summary

| Concept              | Explanation                               |
| -------------------- | ----------------------------------------- |
| **Value Receiver**   | Gets a copy; doesn’t change original      |
| **Pointer Receiver** | Can modify the original struct            |
| **Methods**          | Belong to types, unlike functions         |
| **Interfaces**       | Define behavior via method signatures     |
| **Embedded Structs** | Reuse fields and methods like inheritance |

---
