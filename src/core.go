package main

import (
  "bufio"
  "fmt"
  "os"
  "flag"
)

func defineArgs() {
  // Defines args for the flag package.
  var humArg = flag.String("i", "null", "Accepts human definition of f/w rule.")
}

func parseStdIn() {
  // Parses the user's input.
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    fmt.Println(scanner.Text())
  }
}

func main() {
  parseStdIn()
}