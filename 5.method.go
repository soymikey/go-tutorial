// 6.1 Method Declarations
// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Point is a 2D point
// type Point struct {
// 	X, Y float64
// }

// // Distance calculates the distance from the point to the origin
// func (p Point) Distance() float64 {
// 	return math.Sqrt(p.X*p.X + p.Y*p.Y)
// }

// func main() {
// 	p := Point{3, 4}
// 	fmt.Println(p.Distance()) // Output: 5
// }

// 6.2 Methods with Pointer Receivers
// package main

// import "fmt"

// type Point struct {
//     X, Y float64
// }

// // ScaleBy modifies the receiver
// func (p *Point) ScaleBy(factor float64) {
//     p.X *= factor
//     p.Y *= factor
// }

// func main() {
//     p := Point{1, 2}
//     p.ScaleBy(2)
//     fmt.Printf("%+v\n", p) // Output: {X:2 Y:4}
// }

// 6.3 Composing Types by Struct Embedding
// package main

// import "fmt"

// type Point struct {
// 	X, Y float64
// }

// type ColoredPoint struct {
// 	Point // Embedded Point
// 	Color string
// }

// func (p Point) Distance() float64 { return 0 /* ... */ }

// func main() {
// 	cp := ColoredPoint{
// 		Point: Point{1, 2},
// 		Color: "red",
// 	}
// 	fmt.Println(cp.Distance()) // Calls Point.Distance
// }

// 6.4 Method Values and Expressions
// package main

// import "fmt"

// type Point struct {
// 	X, Y float64
// }

// func (p Point) Add(q Point) Point {
// 	return Point{p.X + q.X, p.Y + q.Y}
// }

// func main() {
// 	p := Point{1, 2}
// 	q := Point{3, 4}

// 	// Method value
// 	add := p.Add
// 	fmt.Println(add(q))

// 	// Method expression
// 	add2 := Point.Add
// 	fmt.Println(add2(p, q))
// }

// 6.5 Encapsulation
// package main

// // Counter is an encapsulated counter type
// type Counter struct {
//     value int // unexported field
// }

// // Increment increases the counter by one
// func (c *Counter) Increment() {
//     c.value++
// }

// // Value returns the current value
// func (c *Counter) Value() int {
//     return c.value
// }

// func main() {
//     c := Counter{}
//     c.Increment()
//     // c.value++ // This would not compile if Counter is in another package
// }

// package bank

// import "fmt"

// // Account represents a bank account
// type Account struct {
// 	balance int // unexported field
// }

// // Deposit adds money to the account
// func (a *Account) Deposit(amount int) error {
// 	if amount <= 0 {
// 		return fmt.Errorf("invalid deposit amount: %d", amount)
// 	}
// 	a.balance += amount
// 	return nil
// }

// // Balance returns the current balance
// func (a *Account) Balance() int {
// 	return a.balance
// }

// // Withdraw removes money from the account
// func (a *Account) Withdraw(amount int) error {
// 	if amount <= 0 {
// 		return fmt.Errorf("invalid withdrawal amount: %d", amount)
// 	}
// 	if amount > a.balance {
// 		return fmt.Errorf("insufficient funds")
// 	}
// 	a.balance -= amount
// 	return nil
// }
