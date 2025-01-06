package main

// 3.1. Integers
// func main() {
// 	// int, int8, int16, int32, int64
// 	var a int = 10
// 	var b int32 = 20
// 	sum := a + int(b) // Type conversion
// 	fmt.Println("Sum:", sum)
// }

// 3.2. Floating-point numbers
// func main() {
// 	var f1 float32 = 3.14      //6-7 decimal digits
// 	var f2 float64 = 2.718     //15 decimal digits
// 	result := f1 + float32(f2) // Type conversion
// 	fmt.Println("Result:", result)
// }

// 3.3. Complex numbers
// complex64 and complex128.
// func main() {
//     c1 := complex(1, 2) // 1 + 2i
//     c2 := complex(3, 4) // 3 + 4i
//     sum := c1 + c2
//     fmt.Println("Sum:", sum)
// }

// 3.4. Booleans
// &&, ||, and !
// func main() {
//     a := true
//     b := false
//     fmt.Println("a AND b:", a && b)
//     fmt.Println("a OR b:", a || b)
//     fmt.Println("NOT a:", !a)
// }

// 3.5. Strings
// 3.5.1 String literals
// func main() {
// 	raw := `This is a raw string literal

// that spans multiple lines.`
// 	interpreted := "This is an interpreted string literal."
// 	fmt.Println(raw)
// 	fmt.Println(interpreted)
// }

// 3.5.2 Unicode
// func main() {
//     s := "Hello, 世界"
//     fmt.Println(s)
// }

// 3.5.3 UTF-8
// func main() {
//     s := "Hello, 世界"
//     for i, r := range s {
//         fmt.Printf("Index: %d, Rune: %c\n", i, r) // h , e , l , l , o , , , 世 , 界
//     }
// }

// 3.5.4 Strings and byte slices

// func main() {
// 	s := "hello"
// 	b := []byte(s)
// 	fmt.Println(b) // [104 101 108 108 111]
// 	s2 := string(b)
// 	fmt.Println(s2) // "hello"
// }

// 3.5.5 Conversion between strings and numbers
// func main() {
// 	s := "123"
// 	n, err := strconv.Atoi(s)
// 	if err == nil {
// 		fmt.Println(n) // 123
// 	}

// 	s2 := strconv.Itoa(n)
// 	fmt.Println(s2) // "123"
// }

// 3.6. Constants

// const Pi = 3.14159

// func main() {
// 	fmt.Println("Pi:", Pi)
// }

// 3.6.1 Constant generator iota

// const (
// 	a = iota // 0
// 	b        // 1
// 	c        // 2
// )

// func main() {
// 	// fmt.Println(a, b, c) 0, 1, 2
// }

// 3.6.2 Untyped constants
// const (
// 	untypedInt    = 42
// 	untypedString = "Hello"
// )

// func main() {
// 	fmt.Println(untypedInt, untypedString)
// }
