// Simple reading and writing to IO in Go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read content of a file
	content, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		return
	}

	// Print out the content
	fmt.Println(string(content))

	// Write (or overwrite a file)
	outPath := "out/out.txt"
	file, err := os.Create(outPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Write a string to the text file
	writeString := "A message goes here..."
	file.WriteString(writeString)

	// Log (optional)
	fmt.Println("\n", writeString, "\nwas written to: ", outPath)
}