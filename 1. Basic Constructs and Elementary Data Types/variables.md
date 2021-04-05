##### Variables

- Are named values.
- All variables are typed values.
- Variables declared within function bodies are called `local variable`.
- Variables declared out of any function body are called `package-level variables` or `global variables`.

---

- **Variable Declaration** 

- `Standard variable declaration forms`: starts with `var` keyword, followed by declared variable name(var name must be an `identifiers`)
- If a value is not passed at the time of variable declaration then it would be initialized to its `zero values`.
- e.g.:
  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var x int
  	var y float32 = 0.1
  	var name string
  
  	fmt.Printf("x - Type: %T, value: %v\n", x, x)
  	fmt.Printf("y - Type: %T, value: %v\n", y, y)
  	fmt.Printf("name - Type: %T, value: %v\n", name, name)
  }
  ```
  ```bash
  x - Type: int, value: 0
  y - Type: float32, value: 0.1
  name - Type: string, value: 
  ```

- ```go var v1, v2 int = 10, 20```: multiple variables can be declared together in a single statement but should be of one type
  
  