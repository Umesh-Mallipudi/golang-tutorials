
---

````markdown
# Go Pointers and Method Receivers – Summary

## 📌 1. Struct Instantiation

```go
u := User{Name: "Alice"}     // Value of type User
pu := &u                     // Pointer to User
pu2 := &User{Name: "Bob"}    // Pointer to new User directly
````

* `u` has type `User`
* `pu` and `pu2` have type `*User`

---

## 🔁 2. & and \* Operators in Go

| Symbol | Name                 | Meaning                           | Example       |
| ------ | -------------------- | --------------------------------- | ------------- |
| `&`    | Address-of operator  | Gets the memory address (pointer) | `p := &x`     |
| `*`    | Dereference operator | Gets the value at a pointer       | `v := *p`     |
| `*T`   | Pointer type         | Declares a pointer to type `T`    | `var p *User` |

---

## 🧠 3. Method Receivers: Value vs Pointer

### Value Receiver

```go
func (u User) UpdateName(name string) {
    u.Name = name
}
```

* **Receives a copy** of the struct.
* Changes do **not** affect the original.
* Safer for small, immutable structs.

### Pointer Receiver

```go
func (u *User) UpdateName(name string) {
    u.Name = name
}
```

* Receives a **pointer to the original struct**.
* Can **modify the original struct**.
* More **efficient** for large structs (only passes 8 bytes).

---

## 🧪 4. Example: Value vs Pointer

```go
type User struct {
    Name string
}

func (u User) SetValue(name string) {
    u.Name = name
}

func (u *User) SetPointer(name string) {
    u.Name = name
}

func main() {
    u := User{Name: "A"}
    u.SetValue("B")
    fmt.Println(u.Name) // Output: A (no change)

    u.SetPointer("C")
    fmt.Println(u.Name) // Output: C (changed)
}
```

---

## 🔁 5. Shared Pointers vs Separate Instances

### Separate Structs

```go
u1 := User{Name: "A"}
u2 := User{Name: "B"}
u1.UpdateName("Z")
// Only u1 is changed
```

### Shared Pointer

```go
shared := &User{Name: "A"}
u1 := shared
u2 := shared
u1.UpdateName("Z")
// Both u1 and u2 point to same object → both are changed
```

---

## 🧾 6. Recap

* Use `*T` in method receivers to modify the original object.
* Use `&T{}` to create a pointer to a new struct.
* `*ptr` gets the actual value from a pointer.
* Passing pointers improves performance by avoiding struct copying.
* Only structs sharing the same pointer reflect changes across variables.

---

