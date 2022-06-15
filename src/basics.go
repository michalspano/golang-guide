package main // Default 'main package'

// Include packages
import (
	"errors"
	"fmt"
	"math"
)

// Create a structure
type person struct {
	name string
	age int
}

func main() {

	// Standard output
	fmt.Println("Hello, World!")

	// User input in GO
	var name string
	fmt.Print("Enter you name: ")
	fmt.Scanln(&name)
	fmt.Println("Your name is: ", name)

	// Variables
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// Conditions
	if x > 6 {
		fmt.Println(".")
	} else if y < 2 {
		fmt.Println("..")
	} else {
		fmt.Println("...")
	}

	// Arrays (fixed size)
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Print(arr)

	// Slice (works just like a list)
	a := []int{1, 2, 3, 4, 5}

	// Append new element (to a slice)
	a = append(a, 15)
	fmt.Println(a)

	// Map a.k.a. dictionary
	dict := make(map[string]int)
	dict["foo"] = 1
	dict["bar"] = 2
	dict["baz"] = 3

	fmt.Println(dict)

	// Access specific elements
	fmt.Println(dict["foo"])

	// Delete elements
	delete(dict, "foo")
	fmt.Println(dict)

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println()
	}

	// While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Iterate over elements in a slice or arr
	myArr := []string{"a", "b", "c"}
	for index, value := range myArr {
		fmt.Println("Index", index, "Value", value)
	}

	// Iterate over keys and values in a map
	myMap := make(map[string]int)
	myMap["a"] = 0
	myMap["b"] = 1
	for key, value := range myMap {
		fmt.Println("Key", key, "Value", value)
	}

	// Custom function
	result := f(10, 20)
	fmt.Println(result)

	// More return values in a custom function
	sqrtResult, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrtResult)
	}

	// Access custom structs
	p := person{name: "Michal", age: 18}
	fmt.Println(p)

	// Specific field of a struct
	fmt.Println(p.age)

	// Pointers (similar as C)
	num1 := 10
	fmt.Println("Initial value: ", num1)

	increment(&num1)
	fmt.Print("Increment value: ", num1, "\n")
}

// My custom function (with a single return)
func f(x int, y int) int {
	return x + y
}

// My second custom function (with more returns)
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("NEGATIVE SQUARE ROOT DOES NOT EXIST")
	}
	return math.Sqrt(x), nil
}

// Function with pointers
func increment(x *int) {
	*x++
}
