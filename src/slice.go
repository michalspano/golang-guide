// Slices in GO
package main

import "fmt"

func main() {

	// Create an empty slice with non-zero length
	sLength := 3
	slice := make([]string, sLength)
	fmt.Println("An empty slice:", slice)

	// Populate the slice
	for i := 0; i < sLength; i++ {
		slice[i] = string(rune(i + 65))  // <- Turn ASCII values to char
	}
	fmt.Println("Non-empty slice: ", slice)

	// Reference by an index
	idx := 1
	fmt.Println("Value of index: '", idx, "' -> ", slice[idx])

	// Append new element(s)
	slice = append(slice, "D", "E")

	// Provide even more elements (in a loop)
	for j := len(slice); j < sLength + 5; j++ {
		slice = append(slice, string(rune(j + 65)))
	}
	fmt.Println(slice)

	// Copy slices
	newSLice := make([]string, len(slice))
	copy(newSLice, slice)
	fmt.Println("Copied slice: ", newSLice)

	// Slice operator (works just like Python ;))
	idx1, idx2 := 1, 3
	sliceOp := slice[idx1:idx2] // idx2 excluded
	fmt.Println(idx1, idx2, " -> ", sliceOp)

	// Initialise a two-dimensional slice
	twoDimLen := 5
	twoDim := make([][]string, twoDimLen)

	// Outer slice
	for i := 0; i < twoDimLen; i++ {
		innerLen := i + 1  // Length of the current inner slice
		twoDim[i] = make([]string, innerLen)

		// Populate inner slice of length `innerLen`
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = string(rune(i + 97))
		}
	}

	fmt.Println("Two-dimensional slice: ", twoDim)
}