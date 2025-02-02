// 7.1 Interfaces as Contracts

// package main

// import "fmt"

// // Stringer is an interface that requires a String method
// type Stringer interface {
// 	String() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// // String implements the Stringer interface
// func (p Person) String() string {
// 	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
// }

// func main() {
// 	var s Stringer = Person{"Alice", 30}
// 	fmt.Println(s.String()) // Output: Alice (30 years old)
// }

// 7.2 Interface Types

// package main

// import "fmt"

// // Abstraction using interface
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// func ReadData(r Reader) {
// 	buf := make([]byte, 100)
// 	n, err := r.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Read", n, "bytes")
// }

// 7.3 Interface Implementation
// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	var r io.Reader = strings.NewReader("Hello, Go!")
// 	buf := make([]byte, 4)
// 	n, _ := r.Read(buf)
// 	fmt.Println(string(buf[:n])) // Output: Hell
// }

// 7.4 Parsing Flags with flag.Value
// package main

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// type CSVFlag []string

// func (c *CSVFlag) String() string {
// 	return fmt.Sprint(*c)
// }

// func (c *CSVFlag) Set(value string) error {
// 	*c = strings.Split(value, ":")
// 	return nil
// }

// func main() {
// 	var list CSVFlag
// 	flag.Var(&list, "list", "Comma-separated list")
// 	flag.Parse()
// 	fmt.Println("List:", list)
// }

// 7.5 Interface Values
// package main

// import "fmt"

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main() {
// 	var i interface{}
// 	i = 42
// 	describe(i) // Output: (42, int)
// 	i = "hello"
// 	describe(i) // Output: (hello, string)
// }

// 7.6 Sorting with sort.Interface
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type Person struct {
// 	Name string
// 	Age  int
// }

// type ByAge []Person

// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func main() {
// 	people := []Person{
// 		{"Bob", 31},
// 		{"Alice", 25},
// 		{"Charlie", 29},
// 	}
// 	sort.Reverse(ByAge(people))
// 	fmt.Println(people)
// }

// 7.7 The http.Handler Interface
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type HelloHandler struct{}

// func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, Go!")
// }

// func main() {
// 	var h HelloHandler
// 	http.ListenAndServe(":8080", h)
// }

// 7.8 The error Interface
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	err := errors.New("an error occurred")
// 	if err != nil {
// 		fmt.Println(err) // Output: an error occurred
// 	}
// }
