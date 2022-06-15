// Passing command line arguments in go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Command line arguments without Prog
	argv := os.Args[1:]

	// Number of command line arguments
	argc := len(argv)
	fmt.Println("Number of command line arguments", argc)

	// Print out all command line arguments
	for i := range argv {
		fmt.Print(argv[i], " ")
	}

	// Print out new line
	fmt.Println()
}