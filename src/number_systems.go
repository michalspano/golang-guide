package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the command-line arguments
	args := os.Args[1:]
	if (len(args) != 1) {
		fmt.Println("Enter a text to read.")
		return
	}

	// Header
	fmt.Println("Converting a text into various number systems.")
	for _, char := range args[0] {

		// Assigning individual values in different number systems
		binVal := fmt.Sprintf("%b", char)
		hexVal := fmt.Sprintf("%x", char)
		octVal := fmt.Sprintf("%o", char)		

		// Print out the values
		fmt.Printf("`%c`: %d | ", char, int(char))
		fmt.Printf("bin: %s, hex: %s, oct: %s\n", binVal, hexVal, octVal)
	}
}