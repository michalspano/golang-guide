// Pointers (oh, everyone's beloved pointers)
// A pinter is hexadecimal address of a variable in memory

package main

import (
	"fmt"
)


// Reset variables value without global reference (using pointers instead)
func resetValue (valueAddress *int, reset int) {
	*valueAddress = reset  // Dereference the hex address and assign new int value
}

// Main function
func main() {

	// Initial var
	i := 1
	fmt.Println("Initial value of i: ", i)

	// Increment var
	i++
	fmt.Println("Just incremented the value of i to: ", i)

	// Reset value using its pointer
	resetValue(&i, 1)
	fmt.Println("`i` has been reset to: ", i)

	// Print variable's pointer value
	fmt.Println("Pointer: ", &i)

	// Pointer of an indexed loop
	for j := 0; j <= 10; j++ {
		fmt.Println("P. ", j, "address: ", &j)
	}

	// Pointers in an array
	arr := [3]string{"A", "B", "C"}
	for i := range arr {
		fmt.Println("Value: ", arr[i], "\nPointer: ", &arr[i])
	}

	// Pointers in a map (each iteration preserves pointer addresses)
	dict := make(map[string]int)

	// Feed data (new pointer values if iteration)
	for i := 0; i < 5; i++ {
		dict[string(rune(i + 65))] = i + 1
	}

	// Show key and value pointers from the map (they remain the same, each iteration updates preserves pointer addresses)
	for key, value := range dict {
		fmt.Println(key, "p: ", &key, value, "p: ", &value)
	}
}