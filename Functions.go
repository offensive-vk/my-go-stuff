package main

import "fmt"

// Function to calculate the sum of a slice
func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    // Slice of integers
    nums := []int{1, 2, 3, 4, 5}

    // Print slice and its sum
    fmt.Println("Numbers:", nums)
    fmt.Println("Sum:", sum(nums))
}
