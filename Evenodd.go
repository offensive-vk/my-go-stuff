package main

import "fmt"

// Function to print even and odd numbers in the given range
func printEvenOdd(start, end int) {
	fmt.Println("Even Numbers:")
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}

	fmt.Println("\nOdd Numbers:")
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var start, end int

	// Take input from the user
	fmt.Print("Enter the start of the range: ")
	fmt.Scanln(&start)
	fmt.Print("Enter the end of the range: ")
	fmt.Scanln(&end)

	// Call the function to print even and odd numbers
	printEvenOdd(start, end)
}
