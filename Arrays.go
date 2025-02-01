package main

import "fmt"

func main() {
    // 1. Fixed-size Array (Classic array)
    var fixedArray [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Fixed-size Array:", fixedArray)

    // 2. Inferred-size Array
    inferredArray := [...]int{10, 20, 30, 40}
    fmt.Println("Inferred-size Array:", inferredArray)

    // 3. Multi-dimensional Array (2D Array)
    var matrix [2][3]int = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println("Multi-dimensional Array:")
    for _, row := range matrix {
        fmt.Println(row)
    }

    // 4. Array of Strings
    var stringArray = [3]string{"Go", "Rust", "Python"}
    fmt.Println("String Array:", stringArray)

    // 5. Boolean Array
    var boolArray = [3]bool{true, false, true}
    fmt.Println("Boolean Array:", boolArray)

    // 6. Default Zero-Initialized Array
    var defaultArray [4]int
    fmt.Println("Default Zero-Initialized Array:", defaultArray)

    // 7. Accessing Elements
    fmt.Println("First Element of Fixed Array:", fixedArray[0])
    fmt.Println("Last Element of String Array:", stringArray[len(stringArray)-1])

    // 8. Modifying an Array
    fixedArray[2] = 99
    fmt.Println("Modified Fixed-size Array:", fixedArray)
}
