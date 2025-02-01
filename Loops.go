package main

import "fmt"

func main() {
    fmt.Println("1. Basic for loop:")
    for i := 1; i <= 5; i++ { // Classic for loop
        fmt.Print(i, " ")
    }
    fmt.Println("\n")

    fmt.Println("2. For loop as while loop:")
    j := 1
    for j <= 5 { // Acts like a while loop
        fmt.Print(j, " ")
        j++
    }
    fmt.Println("\n")

    fmt.Println("3. Infinite loop (breaking it manually):")
    k := 1
    for { // Infinite loop (must use break)
        if k > 5 {
            break
        }
        fmt.Print(k, " ")
        k++
    }
    fmt.Println("\n")

    fmt.Println("4. Range-based loop (iterating over an array):")
    nums := []int{10, 20, 30, 40, 50}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    fmt.Println()

    fmt.Println("5. Loop with continue (skipping even numbers):")
    for m := 1; m <= 10; m++ {
        if m%2 == 0 {
            continue // Skips even numbers
        }
        fmt.Print(m, " ")
    }
    fmt.Println("\n")

    fmt.Println("6. Nested loops (multiplication table):")
    for x := 1; x <= 3; x++ {
        for y := 1; y <= 3; y++ {
            fmt.Printf("%d x %d = %d\t", x, y, x*y)
        }
        fmt.Println()
    }
}
