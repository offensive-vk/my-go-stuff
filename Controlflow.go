package main

import "fmt"

func main() {
    // if-else
    num := 10
    if num%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // for loop
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // switch statement
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Another day")
    }
}
