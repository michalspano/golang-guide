// JSON Support in GO
package main

// Read data to custom structs
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users struct containing all users
type Users struct {
	// Variable name	Variable type	JSON key value (from `encoding/json`)
	Users []User `json:"users"`
}

// User struct with accessible data
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}


// Social struct containing a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}


// Declare the main function
func main() {
	// Open local JSON file and receive possible errors
	PATH := "json_src/src.json"
	jsonFILE, err := os.Open(PATH)

	// Error handling
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully loaded at`", PATH, "`\n")

	// Defer input file and wrap in an error handler
	defer func(jsonFILE *os.File) {
		err := jsonFILE.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFILE)

	// Read to memory
	parsedData, _ := ioutil.ReadAll(jsonFILE)

	// Create an array of users (type os struct)
	var userData Users

	// Unmarshall a.k.a. parse json content to struct and handle error
	Err := json.Unmarshal(parsedData, &userData)
	if Err != nil {
		fmt.Println(Err)
		return
	}

	// Display parsed content
	display(userData)
}

// Function to display parsed data
func display(data Users) {

	// Iterate over all users
	for i := 0; i < len(data.Users); i++ {
		fmt.Println("Name:", data.Users[i].Name)
		fmt.Println("Type:", data.Users[i].Type)

		// Int to String format
		fmt.Println("Age:", strconv.Itoa(data.Users[i].Age))
		fmt.Println("Facebook:", data.Users[i].Social.Facebook)
		fmt.Println("Twitter:",data.Users[i].Social.Twitter + "\n")
	}
}