package main

import (
	"fmt"
	"os"
)

// Create a global variable of type string
var outPath string = "out/sys.txt"

func main() {

	// Assign current path to a variable
	path, err := os.Getwd()
	if err != nil {
        fmt.Println(err)
        return
    }

	// Empty output file
	_, err = os.Create(outPath)
	if err != nil {
        fmt.Println(err)
        return
    }

	// Call printPath function
	printFiles(path)
}

// Print files in path's directory
// Decide whether the file is a directory or a regular file
// recursively call printFiles() to print files in subdirectories
func printFiles(path string) {
	// Open the path's directory and handle error
    dir, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read the directory and handle error
    files, err := dir.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Close the directory and handle error
	closeErr := dir.Close()
	if closeErr != nil {
		return 
	}

    // Loop through the files
    for _, file := range files {

        // Print the file's name to the console
        fmt.Println(file.Name())
		writeFile(outPath, file.Name() + "\n")

        // If the file is a directory, call printFiles()
        if file.IsDir() {
            printFiles(path + "/" + file.Name())
			writeFile(path + "/" + file.Name(), "\n")
        }
    }
}

// Function to write to a file
// Input: file name, content
// Output: none
func writeFile(fileName string, content string) {
    // Open the file and handle error
    file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Write to the file and handle error
    _, err = file.WriteString(content)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Close the file and handle error
    err = file.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}