// Read data from a text file to memory in GO
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Open input file
	f, err := os.Open("foo.txt")

	// Check for input file error
	if err != nil {
		fmt.Println(err)
	}

	// Create a scanner instance
	scanner := bufio.NewScanner(f)

	// Declare an empty slice
	var content []string

	// Iterate over all lines and populate the slice
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	// Iterate over all elements of the slice and print them out
	for i := range content {
		fmt.Println("Line ", i, ": ", content[i])
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Close input file (with error protection)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
}
