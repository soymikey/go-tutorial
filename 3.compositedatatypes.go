package main

// 4.1 Arrays
// func main() {
// 	var a [5]int                           // Array of 5 integers
// 	a[0] = 1                               // Assign a value
// 	b := [3]string{"Go", "Rust", "Python"} // Initialize with values
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// 4.2 Slices
// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	s := a[1:4] // Slice of array
// 	fmt.Println(s)
// }

// 4.2.1 append Function
// func main() {
// 	var s []int
// 	s = append(s, 1, 2, 3, 4, 5)
// 	fmt.Println(s)
// }

// 4.2.2 Slice Modifications
// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	s := a[1:4] //[ 2, 3, 4]
// 	s[0] = 9       // Modifies the underlying array [ 9, 3, 4]
// 	fmt.Println(a) // Now [1, 9, 3, 4, 5]
// }

// 4.3 Maps
// func main() {
// 	m := map[string]int{"apple": 2, "banana": 3}
// 	m["cherry"] = 5
// 	fmt.Println(m)
// 	delete(m, "banana")
// 	fmt.Println(m)
// }

// 4.4 Structs

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	var p Person
// 	p.Name = "Alice"
// 	p.Age = 30
// 	fmt.Println(p)
// }

// 4.4.1 Struct Literals

// type Point struct {
//     X, Y int
// }

// func main() {
//     p := Point{X: 1, Y: 2}
//     fmt.Println(p)
// }

// 4.4.2 Struct Comparison

// import "fmt"

// type Point struct {
// 	X, Y int
// }

// func main() {
// 	p1 := Point{1, 2}
// 	p2 := Point{1, 2}
// 	fmt.Println(p1 == p2) // true
// }

// 4.4.3 Struct Embedding and Anonymous Fields

// import "fmt"

// type Address struct {
// 	City, State string
// }

// type Person struct {
// 	Name string
// 	Age  int
// 	Address
// }

// func main() {
// 	p := Person{
// 		Name:    "Alice",
// 		Age:     30,
// 		Address: Address{City: "New York", State: "NY"},
// 	}
// 	fmt.Println(p.City) // Access embedded field directly
// }

// 4.5 JSON
// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	p := Person{Name: "Alice", Age: 30}
// 	data, _ := json.Marshal(p)
// 	fmt.Println(string(data))

// 	var p2 Person
// 	json.Unmarshal(data, &p2)
// 	fmt.Println(p2)
// }

// 4.6 Text and HTML Templates

// func main() {
// 	const tmpl = `Hello, {{.Name}}!`
// 	t := template.Must(template.New("greeting").Parse(tmpl))
// 	t.Execute(os.Stdout, struct{ Name string }{Name: "World"})
// }

// func main() {
// 	const tmpl = `<h1>Hello, {{.Name}}!</h1>`
// 	t := template.Must(template.New("greeting").Parse(tmpl))
// 	t.Execute(os.Stdout, struct{ Name string }{Name: "World"})
// }
