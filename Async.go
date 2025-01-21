package main

import (
	"fmt"
	"time"
)

// Simulate an async task
func fetchData(resultChan chan string, delay time.Duration) {
	fmt.Println("Fetching data...")
	time.Sleep(delay) // Simulate a long-running operation
	resultChan <- "Data fetched successfully!"
}

func main() {
	// Create a channel to receive the result
	resultChan := make(chan string)

	// Start a goroutine (mimics async task)
	go fetchData(resultChan, 2*time.Second)

	fmt.Println("Doing other work while waiting for the result...")

	// Await the result from the channel (blocks until the result is available)
	result := <-resultChan
	fmt.Println(result)

	fmt.Println("All tasks completed.")
}
