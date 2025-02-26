// package main

// race condition definition: A race condition occurs when two goroutines access a shared variable concurrently
// and at least one of the accesses is a write.

// 8.1 Race Conditions
// var balance int

// func Deposit(amount int) {
// 	balance = balance + amount
// }

// func Balance() int {
// 	return balance
// }

// func runBank() {
// 	//Alice
// 	go func() {
// 		Deposit(200)
// 		fmt.Println("Alice", Balance())
// 	}()
// 	//Bob
// 	go func() {
// 		Deposit(100)
// 	}()
// 	time.Sleep(time.Second)
// 	fmt.Println("Balance", Balance())
// }

// runBank 4 possible results:
// Alice first, then Bob, then Balance 300
// Bob first, then Alice, then Balance 300
// Alice and Bob, then Alice, then Balance 300
// Alice read balance 0, then Bob deposit 100, then Alice deposit 200, then Balance 200

// Three ways to handle race condition:
// 1. Don't modify the variable
// 2. Avoid accessing the variable from multiple goroutines
// 3. Allow only one goroutine to access the variable at a time

// var deposits = make(chan int)
// var balances = make(chan int)

// func Deposit(amount int) {
// 	deposits <- amount
// }

// func Balance() int {
// 	return <-balances
// }

// func teller() {
// 	var balance int
// 	for {
// 		select {
// 		case amount := <-deposits:
// 			balance += amount
// 		case balances <- balance:
// 		}
// 	}
// }

// func runBank() {
// 	go teller()

// 	go func() {
// 		Deposit(200)
// 		fmt.Println("Alice", Balance())
// 	}()
// 	go func() {
// 		Deposit(100)
// 		fmt.Println("Bob", Balance())
// 	}()
// 	time.Sleep(time.Second)
// 	fmt.Println("Balance", Balance())
// }

// 8.2 Mutual Exclusion: sync.Mutex

// const LIMIT = 10000000 // 10 million

// var (
// 	mu       sync.Mutex
// 	balance1 int
// )

// func Bank1Deposit(amount int) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	balance1 += amount
// }

// func Bank1Balance() int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	return balance1
// }

// func bank1() {
// 	go Bank1Deposit(200) // Write operation
// 	go Bank1Deposit(100) // Write operation

// 	// Add timing measurement
// 	start := time.Now()
// 	for i := 0; i < LIMIT; i++ {
// 		go Bank1Balance() // Many more concurrent reads
// 	}
// 	time.Sleep(time.Second)
// 	fmt.Printf("Bank1 (Mutex) took: %v\n", time.Since(start))
// 	fmt.Println("Final Balance:", Bank1Balance())
// }

// // 8.3 Read-Write Mutex: sync.RWMutex

// var (
// 	rwmu     sync.RWMutex
// 	balance2 int
// )

// func Bank2Deposit(amount int) {
// 	rwmu.Lock()
// 	defer rwmu.Unlock()
// 	balance2 += amount
// }

// func Bank2Balance() int {
// 	rwmu.RLock()
// 	defer rwmu.RUnlock()
// 	return balance2
// }

// func bank2() {

// 	go Bank2Deposit(200) // Write operation
// 	go Bank2Deposit(100) // Write operation

// 	// Add timing measurement
// 	start := time.Now()
// 	for i := 0; i < LIMIT; i++ {
// 		go Bank2Balance() // Many more concurrent reads
// 	}
// 	time.Sleep(time.Second)
// 	fmt.Printf("Bank2 (RWMutex) took: %v\n", time.Since(start))
// 	fmt.Println("Final Balance:", Bank2Balance())
// }

// func main() {
// 	fmt.Println("Testing Mutex (bank1):")
// 	bank1()

// 	fmt.Println("\nTesting RWMutex (bank2):")
// 	bank2()
// }
