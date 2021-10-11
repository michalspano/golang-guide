package main

import "fmt"


// MAX Defines a constant number of max alphabetical chars
const MAX = 26

func main() {

	// Pass user input
	var Range int
	fmt.Print("Enter desired range: ")
	fmt.Scanln(&Range)

	// Check valid user input
	if !validateRange(Range) {
		panic("Invalid input range.")
	}

	// Create a map
	myMap := make(map[string]int)

	// Assign elements to a map
	for i := 0; i < Range; i++ {
		myMap[string(rune(i+97))] = i + 1
	}

	// Read from a map
	readMap(myMap)
}

// Create a function to check correct input value (return bool)
func validateRange(x int) bool {
	if x > MAX || x < 0 {
		return false
	} else {
		return true
	}
}

// Create a function to read from a map
func readMap(m map[string]int) {
	for key, value := range m {
		fmt.Println("Key:", key, "\nvalue:", value)
	}
}