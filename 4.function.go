// 5.1 Function Declarations
// package main

// import "fmt"

// // Function declaration with parameters and return type
// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(3, 5)
// 	fmt.Println("Sum:", result) // Output: Sum: 8
// }

// 5.2 Recursion
// package main

// import "fmt"

// // Recursive function to compute factorial
// func factorial(n int) int {
// 	if n == 0 {
// 		return 1 // Base case
// 	}
// 	return n * factorial(n-1) // Recursive call  // 5 * 4 * 3 * 2 * 1
// }

// func main() {
// 	fmt.Println("Factorial of 5:", factorial(5)) // Output: 120
// }

// 5.3 Multiple Return Values
// package main

// import "fmt"

// // Function that returns two values
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result) // Output: Result: 5
// }

// 5.4 Errors
// 5.4.1 Error Handling Strategies
// 5.4.1.a Propagating Errors
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result)
// }

// 5.4.1.b Logging and Continuing
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func readFile(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Proceed with file reading
// 	fmt.Println("File opened successfully")
// }

// func main() {
// 	readFile("nonexistent.txt")
// 	fmt.Println("Continuing execution...")
// }

// 5.4.1.c Exiting the Program
// package main

// import (
//     "fmt"
//     "os"
// )

// func checkError(err error) {
//     if err != nil {
//         fmt.Fprintln(os.Stderr, "Critical error:", err)
//         os.Exit(1)
//     }
// }

// func main() {
//     file, err := os.Open("nonexistent.txt")
//     checkError(err)
//     defer file.Close()

//     // Proceed with file reading
//     fmt.Println("File opened successfully")
// }
// 5.4.1.d Recovering from Panic
// package main

// import "fmt"

// func mightPanic() {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Recovered from panic:", r)
//         }
//     }()
//     panic("Something went wrong!")
// }

// func main() {
//     fmt.Println("Start")
//     mightPanic()
//     fmt.Println("End")
// }

// 5.4.1.e Retry operations

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // mockOperation simulates an operation that might fail.
// // It randomly fails to demonstrate the retry mechanism.
// func mockOperation() error {
// 	if rand.Intn(3) == 0 { // Randomly succeed or fail
// 		return nil
// 	}
// 	return errors.New("operation failed")
// }

// // retryOperation attempts to perform an operation with retries.
// func retryOperation(maxRetries int) error {
// 	var err error
// 	for i := 0; i < maxRetries; i++ {
// 		err = mockOperation()
// 		if err == nil {
// 			return nil
// 		}
// 		fmt.Printf("Attempt %d failed, retrying...\n", i+1)
// 		time.Sleep(500 * time.Millisecond) // Wait before retrying
// 	}
// 	return fmt.Errorf("operation failed after %d attempts: %w", maxRetries, err)
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

// 	err := retryOperation(5) // Try up to 5 times
// 	if err != nil {
// 		fmt.Println("Final error:", err)
// 	} else {
// 		fmt.Println("Operation succeeded")
// 	}
// }

// 5.5 Function Values
// package main

// import "fmt"

// // Function that adds two integers
// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     // Assign function to a variable
//     op := add
//     fmt.Println("Sum:", op(2, 3)) // Output: Sum: 5
// }

// 5.6 Anonymous Functions
// package main

// import "fmt"

// func main() {
// 	// Define and immediately invoke an anonymous function
// 	func(message string) {
// 		fmt.Println(message)
// 	}("Hello, Go!") // Output: Hello, Go!
// }

// 5.6.1 Anonymous Functions (closure)
// package main

// import "fmt"

// // Regular function that takes 'count' as an argument
// func increment(count int) int {
// 	count++
// 	return count
// }

// func main() {
// 	count := 0

// 	fmt.Println(increment(count)) // Output: 1
// 	fmt.Println(increment(count)) // Output: 1
// 	fmt.Println(increment(count)) // Output: 1
// }

// package main

// import "fmt"

// func main() {
// 	count := 0

// 	// Anonymous function (closure) that captures 'count'
// 	increment := func() int {
// 		count++
// 		return count
// 	}

// 	fmt.Println(increment()) // Output: 1
// 	fmt.Println(increment()) // Output: 2
// 	fmt.Println(increment()) // Output: 3
// }

// 5.7 Variadic Functions
// package main

// import "fmt"

// // Variadic function that sums integers
// func sum(nums ...int) int {
//     total := 0
//     for _, num := range nums {
//         total += num
//     }
//     return total
// }

//	func main() {
//	    fmt.Println("Total:", sum(1, 2, 3, 4, 5)) // Output: Total: 15
//	}
//
// 5.8 Deferred Function Calls
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start")
// 	defer fmt.Println("Deferred 1")
// 	defer fmt.Println("Deferred 2")
// 	fmt.Println("End")
// }

// 5.9 Panic
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Before panic")
// 	panic("Something went wrong!")
// 	fmt.Println("After panic") // This line is not executed
// }

// 5.10 Recover
// package main

// import "fmt"

// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 		}
// 	}()
// 	fmt.Println("Before panic")
// 	panic("Unexpected error!")
// 	fmt.Println("After panic") // This line is not executed
// }
