package main

import (
  "fmt"
  "flag"
  "strings"
  "errors"
)

// Parses the human's plaintext input.
func parseStdIn(humArgPtr *string)(int, error) {
  acceptedVerbs := map[string]bool {
      "permit": true,
      "block": true,
  }
  if acceptedVerbs[*humArgPtr] {
      fmt.Println("Accepted Verbs Exist")
  }
  else {
    // Fail Fast
    return -1, errors.New("Statement does not include an accepted verb.")
  }
}

func main() {
  // Retrieve User Input
  var humArgPtr = flag.String("i", "null", "Accepts human definition of f/w rule.")
  flag.Parse()

  parseStdIn(humArgPtr)
}