// Closures in Go
package main

// Import libs
import (
	"fmt"
	"os"
	"strconv"
)


// Create a function to parse closure sequence
func customClosure() func() int {
	i := 0

	// Return incremented value (until new function call)
	return func() int {
		i++
		return i
	}
}


func main() {

	// Import range from command line
	argv := os.Args[1:]

	// Check valid command-line arguments
	if len(argv) != 1 {
		panic("Incorrect number of command-line arguments.")
	}

	// Assign range and check type conversion errors
	closureRange, err := strconv.Atoi(argv[0]); if err != nil {
		panic(err)
	}

	// Initial closure call
	closureCall := customClosure()


	// Print closure R times
	for j := 0; j < closureRange; j++ {

		// Reset closure (a.k.a. iterate from initial index)
		if j % 5 == 0 {
			fmt.Println("Closure was reset.")
			closureCall = customClosure()
		}

		// Print current closure value
		fmt.Println("Closure call no. ", closureCall())
	}
}