Great! In Go, you can get **user input** using the `fmt` or `bufio` package.

---

### ✅ **1. Using `fmt.Scan` (Simple and Direct)**

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // reads input until the first space or newline
    fmt.Println("Hello,", name)
}
```

**🟡 Note**: `fmt.Scan` reads **only one word**. For full lines, use `bufio`.

---

### ✅ **2. Using `bufio.NewReader` (For full line input)**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your full name: ")
    input, _ := reader.ReadString('\n')  // reads until newline
    input = strings.TrimSpace(input)     // removes newline character
    fmt.Println("Welcome,", input)
}
```

---

### ✅ **3. Example with Number Input**

```go
package main

import (
    "fmt"
)

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Println("You are", age, "years old.")
}
```

---

### 🧠 Summary:

| Method            | Reads         | Use for                                    |
| ----------------- | ------------- | ------------------------------------------ |
| `fmt.Scan`        | One word      | Simple inputs (e.g., int, one-word string) |
| `fmt.Scanln`      | Until newline | Slightly better than `Scan`                |
| `bufio.NewReader` | Full line     | Full name, sentence input                  |


