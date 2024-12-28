// 1. Names
// CamelCase Naming Convention:
var myVariable string
const myConstant string = "Hello, World!"
type myType int
func myFunction() {}

// Package Imports:
import ("fmt","strings")

// 2. Declarations
// Variable Declarations:
var x int 
var y=10 // Type Inference
var z int = 10

// Constant Declarations:
const Pi = 3.14

// Type Declarations:
type Celsius float64

// 3. Variables
// Variable Declaration and Initialization:
var x int = 10
var a= 2 // Type Inference
c:=3 // Short Variable Declaration

// Zero Value of Variables:
var x int // 0
var y float64 // 0.0
var z string // ""


// 4. Assignments
// Simple Assignment:
var x int
x = 10

// Compound Assignment:
x+=5
x-=5

// Short Variable Declaration Assignment:
x:=10

// 5. Type Declarations
// Type Alias:
type myInt int

// New Type Definition:
type Celsius float64
type Fahrenheit float64

x:=Celsius(10.0)
y:=Fahrenheit(10.0)

// 6. Packages and Files
// Package Declaration:
package main

// Importing Packages:
import( "fmt","strings")

// Package Initialization:
func init() {
	fmt.Println("Hello, World!")
}

// 7. Scope
// Package-Level Scope:
var X int
func main() {
	fmt.Println(x)
}
// Function-Level Scope:
func f() {
    localVariable := 1
    if localVariable == 1 {
        fmt.Println("Local variable is 1")
    }
}

