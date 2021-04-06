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
  	var x1 int = 10
  	var y float32
  	var y1 float32 = 10.01
  	var lang string
  	var lang1 string = "Golang"
  
  	fmt.Printf("x - Type: %T, value: %v\n", x, x)
  	fmt.Printf("x1 - Type: %T, value: %v\n", x1, x1)
  	fmt.Printf("y - Type: %T, value: %v\n", y, y)
  	fmt.Printf("y1 - Type: %T, value: %v\n", y1, y1)
  	fmt.Printf("lang - Type: %T, value: %v\n", lang, lang)
  	fmt.Printf("lang1 - Type: %T, value: %v\n", lang1, lang1)
  }
  ```
  ```bash
  x - Type: int, value: 0
  x1 - Type: int, value: 10
  y - Type: float32, value: 0
  y1 - Type: float32, value: 10.01
  lang - Type: string, value: 
  lang1 - Type: string, value: Golang
  ```

- ```go var v1, v2 int = 10, 20```: multiple variables can be declared together in a single statement but should be of one type
  
  