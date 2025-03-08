package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Show message
	fmt.Println("🛑 WARNING: This program will allocate a large amount of RAM!")
	fmt.Print("💾 Enter the amount of RAM to consume (in MB): ")

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim newline and convert to integer
	input = strings.TrimSpace(input)
	ramMB, err := strconv.Atoi(input)
	if err != nil || ramMB <= 0 {
		fmt.Println("❌ Invalid input. Please enter a positive integer.")
		return
	}

	// Convert MB to Bytes
	ramBytes := ramMB * 1024 * 1024
	fmt.Printf("🚀 Allocating %d MB of RAM...\n", ramMB)

	// Allocate memory
	memory := make([]byte, ramBytes)

	// Fill memory to prevent optimizations
	for i := range memory {
		memory[i] = 1
	}

	fmt.Println("✅ Memory allocation complete. Press ENTER to release memory and exit...")

	// Keep the program running until the user presses ENTER
	reader.ReadString('\n')

	// Memory is freed when the program exits
	fmt.Println("🔄 Releasing memory... Done!")
}
