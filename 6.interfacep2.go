// 7.10 Type Assertions
// package main

// import "fmt"

// func main() {
// 	var i interface{} = "hello"
// 	s := i.(string)
// 	fmt.Println(s) // Output: hello

// 	// Safe type assertion
// 	s, ok := i.(string)
// 	fmt.Println(s, ok) // Output: hello true

// 	f, ok := i.(float64)
// 	fmt.Println(f, ok) // Output: 0 false
// }

// 7.11 Using Type Assertions to Identify Errors
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	_, err := os.Open("nonexistent_file.txt")
// 	if err != nil {
// 		if pathError, ok := err.(*os.PathError); ok {
// 			fmt.Printf("Failed at path: %s\n", pathError.Path)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}
// }

// 7.12 Querying Behaviors with Interface Type Assertions
// package main

// import "fmt"

// type Stringer interface {
// 	String() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
// }

// func main() {
// 	var i interface{} = Person{"Alice", 30}
// 	if s, ok := i.(Stringer); ok {
// 		fmt.Println(s.String()) // Output: Alice (30 years old)
// 	}
// }

// 7.13 Type Switches
// package main

// import "fmt"

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }
// func main() {
// 	do(21)      // Output: Twice 21 is 42
// 	do("hello") // Output: "hello" is 5 bytes long
// 	do(true)    // Output: I don't know about type bool!
// }
