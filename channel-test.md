```go
package main

import (
	"fmt"
	"time"
)


func greet(c chan string) {
	name := <-c
	fmt.Println("Hello", name)
}

func greetUntilQuit(c chan string, q chan int) {
	for {
		select {
		case name := <-c:
			fmt.Println("Hello", name)
		case <-q:
			fmt.Println("Quitting...")
			return
		}
	}
}

func main() {
	c := make(chan string)
	q := make(chan int)
	go greetUntilQuit(c, q)

	c <- "Alice"
	c <- "Bob"
	q <- 0 // Send a signal to quit
	time.Sleep(1 * time.Second) // Wait for a moment to ensure the goroutine has time to process
}```