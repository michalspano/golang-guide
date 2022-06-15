//usr/local/go/bin/go run $0 $@ ; exit

package main

// Defined imports
import (
  "fmt"
  "math/rand"
  "time"
  "os"
  "strconv"
  "unicode"
)

// Defined constants
const ( 
  // FIXME: quite obsolete; avoid using ACII codes with a gap betwen 90 a 97
  letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" 
  OUT       = "exp/out.csv"
  OUT_COUNT = 15
)

func main() {
  argv := os.Args

  // Ensure usage
  if len(argv) != 2 {
    fmt.Println("Usage: ./csv.go <string_size>")
    return
  }
  
  // Reset the seed
  rand.Seed(time.Now().UnixNano())
  
  // Parse the command-line argument and catch errors
  stringSize, err := strconv.Atoi(argv[1])
  if err != nil {
    fmt.Println("Error: string size must be an integer")
    return
  }
  
  // Expect positive integer
  if stringSize < 1 {
    fmt.Println("Error: string size must be greater than 0")
    return
  }
  
  // Open the output file and catch errors
  file, err := os.Create(OUT)
  if err != nil {
    fmt.Println("Error: could not create file")
    return
  }
  
  // Write initial header of the raw CSV file
  file.WriteString("name;age\n")
  
  // Generate random name - age pairs in a loop and write them to CSV file
  for i := 0; i < OUT_COUNT; i++ {
    randomName := generateRandomString(stringSize)
    randomAge := rand.Intn(100 - 10) + 10 // Random age - <10, 100>
  
    // Parse to a string format with `;` as a delimiter
    toWrite := fmt.Sprintf("%s;%d\n", randomName, randomAge)
    file.WriteString(toWrite)

  }
}

// Function to generate a random string from a given range 
func generateRandomString(size int) string {
  // Random size - <5, size>
  s := rand.Intn(size - 5) + 5
  
  // Populate the buffer
  buff := make([]rune, s)
  for i := range buff {
    buff[i] = rune(letters[rand.Intn(len(letters))])
  }
  
  // Expect a name starting with a capital letter
  buff[0] = unicode.ToUpper(buff[0])

  return string(buff)
}
