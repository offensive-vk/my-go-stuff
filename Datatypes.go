package main

import "fmt"

func main() {
    // Boolean
    var isActive bool = true

    // Integer types
    var intVal int = 42
    var int8Val int8 = -128
    var uintVal uint = 400

    // Floating-point types
    var floatVal float32 = 3.14
    var doubleVal float64 = 2.7182818284

    // String
    var message string = "Hello, Go!"

    // Byte (alias for uint8)
    var char byte = 'A'

    // Rune (alias for int32)
    var unicode rune = '\u2602' // Unicode for ☂

    // Array
    var arr = [3]int{1, 2, 3}

    // Slice
    slice := []string{"Go", "Python", "JavaScript"}

    // Map
    languages := map[string]string{"Go": "Google", "Python": "Python Software Foundation"}

    // Print all data types
    fmt.Println("Boolean:", isActive)
    fmt.Println("Integer:", intVal, ", int8:", int8Val, ", uint:", uintVal)
    fmt.Println("Float:", floatVal, ", Double:", doubleVal)
    fmt.Println("String:", message)
    fmt.Printf("Byte: %c, Rune: %c\n", char, unicode)
    fmt.Println("Array:", arr)
    fmt.Println("Slice:", slice)
    fmt.Println("Map:", languages)
}
