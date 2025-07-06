You're welcome! 🙌 Now let’s explore **loops in Go**, which are both simple and powerful.

---

## 🔁 Go Loops Overview

Go only has **one looping keyword**: `for`.
It’s super flexible and covers:

* `while` loops
* classic `for` loops
* infinite loops
* range-based loops

---

## ✅ Basic `for` Loop (like C-style)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

* Initialization: `i := 0`
* Condition: `i < 5`
* Post-step: `i++`

---

## 🔁 While-Style Loop

Go doesn’t have a `while` keyword, but you can do:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

> Exactly like a `while` loop in other languages.

---

## 🔁 Infinite Loop

```go
for {
    fmt.Println("Looping forever...")
}
```

* Use this with a `break` condition inside.
* Common for servers, background tasks, etc.

---

## 🔁 Using `break` and `continue`

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // skip this iteration
    }
    if i == 5 {
        break // exit the loop
    }
    fmt.Println(i)
}
```

---

## 🔁 Looping with `range`

Used for iterating over slices, arrays, maps, strings, etc.

### Example with Slice:

```go
nums := []int{10, 20, 30}
for i, val := range nums {
    fmt.Println("Index:", i, "Value:", val)
}
```

### Example with Map:

```go
m := map[string]int{"a": 1, "b": 2}
for key, val := range m {
    fmt.Println("Key:", key, "Value:", val)
}
```

### Example with String (Unicode-safe!):

```go
for i, char := range "go语言" {
    fmt.Printf("Index: %d, Char: %c\n", i, char)
}
```

---

## 🧠 Summary Table

| Loop Type            | Example                   |
| -------------------- | ------------------------- |
| Classic `for`        | `for i := 0; i < 10; i++` |
| While-style          | `for i < 10`              |
| Infinite loop        | `for {}`                  |
| Range over slice     | `for i, v := range slice` |
| Range over map       | `for k, v := range map`   |
| `break` & `continue` | Control loop flow         |

---

Would you like a sample Go program that uses different types of loops together?
