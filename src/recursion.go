// Recursion in Go
package main

import "fmt"

const max = 5

func main() {
	// Call default recursive function with 0 index
	initialIndex := 0
	recursiveFunction(initialIndex)

	// Initial integers
	num1, num2 := 3, 5

	// Print out computed values
	fmt.Printf("\nFactorial of %d: %d\nFactorial of %d: %d\n", num1, factorial(num1), num2, factorial(num2))

}

// Define a default recursive function
func recursiveFunction(x int) int {
	if x < max {
		fmt.Printf("Recursive call number: %d\n", x + 1)
		return recursiveFunction(x + 1)
	}
	return 0
}

// Recursive factorial function
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}