### Go Variables & Naming — Quick Notes

🔹 Variable Declaration:
- `var x int = 10` → typed
- `x := 10` → short form (walrus), type inferred, used in functions

🔹 Constants:
- `const Pi = 3.14` → fixed value, no `:=`

🔹 Export Rules:
- Capital letter → Exported (public)
- Small letter → Unexported (private)

🔹 Naming Style:
- camelCase → local vars & funcs
- PascalCase → exported names
- Avoid snake_case

💡 Tip: Use `:=` when writing quick function logic. Use `var` for clarity or zero-value