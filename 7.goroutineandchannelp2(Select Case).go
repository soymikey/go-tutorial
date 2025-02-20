// package main

// import (
// 	"fmt"
// 	"time"
// )

// // 8.7 Using select for Multiplexing

// Basic syntax
// select {
// case <-ch1:
//     // Handle received signal from ch1
// case data := <-ch2:
//     // Handle data received from ch2
// case ch3 <- someData:
//     // Send someData to ch3
// default:
//     // Execute when no other cases are ready
// }

// func demo() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	ch3 := make(chan string, 1) // Create buffered channel with capacity 1
// 	ready := make(chan bool)    // Channel for synchronization signal

// 	go func() {
// 		ready <- true       // Send ready signal to main goroutine
// 		someData := "Hello" // Data to be sent
// 		select {
// 		case <-ch1:
// 			// Case 1: Receive from ch1 (signal only)
// 			fmt.Println("Received signal from ch1")
// 		case receivedData := <-ch2:
// 			// Case 2: Receive from ch2 (with data)
// 			fmt.Println("Received from ch2:", receivedData)
// 		case ch3 <- someData:
// 			// Case 3: Send data to ch3
// 			fmt.Println("Sent data to ch3")
// 		default:
// 			// Execute if no channel operations are ready
// 			fmt.Println("No channel operations were ready")
// 		}
// 	}()

// 	<-ready                     // Wait for goroutine to be ready
// 	ch2 <- "Hello"              // Send data to ch2
// 	time.Sleep(1 * time.Second) // Keep program alive for 1 more seconds
// }

// // 1. Timeout control
// func timeoutControl() {
// 	// Create a channel for receiving results
// 	resultChan := make(chan string)

// 	// Start a goroutine that simulates a time-consuming task
// 	go func() {
// 		// Simulate work that takes 3 seconds
// 		time.Sleep(1 * time.Second)
// 		resultChan <- "Task completed"
// 	}()

// 	// Use select for timeout control
// 	select {
// 	case result := <-resultChan:
// 		// This case executes if we receive data from resultChan
// 		fmt.Println("Received result:", result)

// 	case <-time.After(2 * time.Second):
// 		// This case executes if 2 seconds pass without receiving data
// 		fmt.Println("Timeout: operation took too long")
// 	}

// 	fmt.Println("Program finished")
// }

// // 2. Non-blocking channel operations
// func nonBlocking() {
// 	// Create a channel with no buffer
// 	dataChan := make(chan string)

// 	// Start a goroutine that sends data after 2 seconds
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		dataChan <- "Hello from goroutine!"
// 	}()

// 	// Try to receive data from channel multiple times
// 	for i := 0; i < 3; i++ {
// 		select {
// 		case data := <-dataChan:
// 			// This case executes when data is available
// 			fmt.Printf("Received data: %s\n", data)

// 		default:
// 			// This case executes immediately if channel is not ready
// 			fmt.Printf("No data available, attempt %d\n", i+1)
// 		}

// 		// Wait for 1 second before next attempt
// 		time.Sleep(1 * time.Second)
// 	}
// }

// // 3. Graceful shutdown
// func worker(id int, done chan bool) {
// 	fmt.Printf("Worker %d starting\n", id)

// 	for {
// 		select {
// 		case <-done:
// 			// Received shutdown signal
// 			fmt.Printf("Worker %d shutting down\n", id)
// 			return
// 		default:
// 			// Simulate some work
// 			fmt.Printf("Worker %d working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func runWorker() {
// 	// Create done channel for shutdown signaling
// 	done := make(chan bool)
// 	go worker(1, done)
// 	time.Sleep(3 * time.Second)
// 	close(done)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main: program finished")
// 	// // Start multiple workers
// 	// for i := 1; i <= 3; i++ {
// 	// 	go worker(i, done)
// 	// }

// 	// // Let workers run for 3 seconds
// 	// fmt.Println("Main: letting workers run for 3 seconds...")
// 	// time.Sleep(3 * time.Second)

// 	// // Signal shutdown to all workers
// 	// fmt.Println("Main: sending shutdown signal to all workers...")
// 	// close(done)

// 	// // Give workers time to finish
// 	// fmt.Println("Main: waiting for workers to shutdown...")
// 	// time.Sleep(1 * time.Second)
// 	// fmt.Println("Main: program finished")

// }

// // 4. Merging multiple channels
// func merge(ch1, ch2 <-chan int) <-chan int {
// 	merged := make(chan int)
// 	go func() {
// 		defer close(merged)
// 		for {
// 			select {
// 			case v := <-ch1:
// 				merged <- v
// 			case v := <-ch2:
// 				merged <- v
// 			}
// 		}
// 	}()
// 	return merged
// }

// // 5. Heartbeat monitoring
// func heartbeat(done chan bool) {
// 	ticker := time.NewTicker(1 * time.Second)
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case <-ticker.C:
// 			// Send heartbeat signal
// 			fmt.Println("ping")
// 		}
// 	}
// }

// func runHeartbeat() {
// 	done := make(chan bool)
// 	go heartbeat(done)
// 	time.Sleep(3 * time.Second)
// 	close(done)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main: program finished")
// }

// func main() {
// 	runHeartbeat()
// }
