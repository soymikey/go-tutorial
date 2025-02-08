// p1
// 8.1 Goroutines
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world") // Start a new goroutine
// 	say("hello")    // Run in the main goroutine
// }

// 8.2 Example: Concurrent Clock Server
// package main

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		go handleConn(conn) // Handle each connection in a new goroutine
// 	}
// }

// 8.3 Example: Concurrent Echo Server
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"strings"
// 	"time"
// )

// func echo(c net.Conn, shout string, delay time.Duration) {
// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", shout)
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
// }

// func handleConn(c net.Conn) {
// 	input := bufio.NewScanner(c)
// 	for input.Scan() {
// 		go echo(c, input.Text(), 1*time.Second) // Handle each input in a new goroutine
// 	}
// 	c.Close()
// }

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		go handleConn(conn) // Handle each connection in a new goroutine
// 	}
// }

// 8.4 Channels
// 8.4.1 Unbuffered Channels
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int) // Create an unbuffered channel
// 	go func() {
// 		fmt.Println("new goroutine")
// 		ch <- 42 // Send a value into the channel
// 	}()
// 	fmt.Println(<-ch) // Receive a value from the channel
// }

// 8.4.2 Pipelines
// package main

// import "fmt"

// data ->naturals->squares-> calculated data
// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	// Generate numbers
// 	go func() {
// 		for x := 0; x < 10; x++ {
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()

// 	// Square numbers
// 	go func() {
// 		for x := range naturals {
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()

// 	// Print squared numbers
// 	for x := range squares {
// 		fmt.Println(x)
// 	}
// }

// 8.4.3 Unidirectional Channel Types

// package main

// import "fmt"

// func counter(out chan<- int) {
// 	for x := 0; x < 10; x++ {
// 		out <- x
// 	}
// 	close(out)
// }

// func squarer(in <-chan int, out chan<- int) {
// 	for x := range in {
// 		out <- x * x
// 	}
// 	close(out)
// }

// func printer(in <-chan int) {
// 	for x := range in {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	go counter(naturals)
// 	go squarer(naturals, squares)
// 	printer(squares)
// }

// 8.4.4 Buffered Channels
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 2) // Create a buffered channel with capacity 2
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch) // Output: 1
// 	fmt.Println(<-ch) // Output: 2
// }

