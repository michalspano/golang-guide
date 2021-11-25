package main

// Handle import and constants
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const (
    defPath string = "/Users/x/foo"
    keyword string = "TODO"
)

// Create the main function
func main() {
    data := detectAllFiles(defPath)
    if data == nil {
        fmt.Printf("No %ss found\n", keyword)
        return
    }

    // Status message
    fmt.Printf("Found %d %ss.\n", len(data), keyword)

    // Iterate over parsed data
    for i := range data {
        for j := range data[i] {
            fmt.Println(data[i][j])
        }
    }
}

// Detect all files from the directory at PATH
func detectAllFiles(PATH string) [][] string{
    // Handle the default path
    path, err := os.Open(PATH)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    // Close the path
    defer func(path *os.File) {
        _ = path.Close()
    }(path)

    // Create the scanner to read the directory
    files, err := path.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    // Create a 2-dim slice of strings
    var temp [][]string = nil

    // Iterate over the files
    for _, file := range files {
        // Ignore if itself
        if file.Name() == "foo.go" {
            continue
        }

        // Read content from the file
        parsedFile := readContentOfFile(PATH + "/" + file.Name())
        if parsedFile != nil {
            temp = append(temp, parsedFile)
        }
    }
    // Return the 2-dim slice
    return temp
}

// Create a function to load content of a text file
func readContentOfFile(PATH string) []string {
    // Open the content of the file
    file, err := os.Open(PATH)
    if err != nil {
        os.Exit(1)
    }

    // Defer - ignore the error
    defer func(path *os.File) {
        _ = path.Close()
    }(file)

    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // Current TODO's found
    var TODOS []string = nil

    // Iterate over the file
    i := 1
    for scanner.Scan() {
        // Detect the column number of the keyword from the current line
        columnNumber := strings.Index(strings.ToUpper(scanner.Text()), keyword)
        if strings.Contains(strings.ToUpper(scanner.Text()), keyword) {
            // Create 'msg' slice to the message after the keyword
            msg := strings.TrimSpace(scanner.Text()[columnNumber+len(keyword) + 1:])
            // Format TODO content
            content := fmt.Sprintf("%s:%d:%d:%s", PATH, i, columnNumber, msg)
            // Populate slice
            TODOS = append(TODOS, content)
            continue
        }
        i++
    }
    // Return TODOs if not nil
    if TODOS == nil {
        return nil
    }
    return TODOS
}
