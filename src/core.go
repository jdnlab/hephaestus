package main

import (
  "fmt"
  "flag"
  "errors"
  "strings"
)

// Parses the human's plaintext input.
// It also validates whether plaintext is valid.
func parseStdIn(humArgPtr *string)(int, error) {
  splitArr := strings.Split(*humArgPtr, " ")

  acceptedVerbs := map[string]bool {
    "permit": true,
    "block": true,
  }

  acceptedProto := map[string]bool {
    "ssh": true,
    "icmp": true,
    "http": true,
  }

  if acceptedVerbs[strings.ToLower(splitArr[0])] {
      fmt.Println("verb validity check: PASSED")
  } else {
    fmt.Println("verb validity check: FAILED")
    return -1, errors.New("statement contains invalid verb")
  }
  if acceptedProto[strings.ToLower(splitArr[1])] {
      fmt.Println("Accepted Protocols Exist")
  } else {
    fmt.Println("protocol validity check: FAILED")
    return -1, errors.New("statement contains invalid protocol")
  }
  fmt.Println("valid statement")
  return 0, nil
}

func main() {
  // Defines options for the human
  var humArgPtr = flag.String("i", "null", "Accepts human definition of f/w rule.")
  flag.Parse()

  parseStdIn(humArgPtr)
}