In Go, **type conversion** means converting a value from one type to another explicitly. Go does **not support implicit type conversion**, so you have to do it manually.

### 🔧 Syntax:

```go
T(v)  // where T is the target type, and v is the value
```

### ✅ Example (Simple types):

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b float64 = float64(a) // convert int to float64

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

### ✅ Example (String to int):

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"
    i, err := strconv.Atoi(s) // convert string to int
    if err != nil {
        fmt.Println("Conversion error:", err)
    } else {
        fmt.Println("Converted int:", i)
    }
}
```

### ✅ Example (Int to string):

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := 42
    s := strconv.Itoa(i) // convert int to string
    fmt.Println("Converted string:", s)
}
```
