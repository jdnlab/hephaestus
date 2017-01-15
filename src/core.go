package main

import (
  "fmt"
  "flag"
  "errors"
  "strings"
  "bytes"
)

// Parses the human's plaintext input.
// It also validates whether plaintext is valid.
func parseStdIn(humArgPtr *string)(int, error) {
  splitArr := strings.Split(*humArgPtr, " ")

  acceptedVerbs := map[string]bool {
    "permit": true,
    "deny": true,
  }

  acceptedProto := map[string]bool {
    "ssh": true,
    "icmp": true,
    "http": true,
  }

  if acceptedVerbs[strings.ToLower(splitArr[0])] {
    fmt.Println("debug verb validity check: PASSED")
  } else {
    fmt.Println("verb validity check: FAILED")
    return -1, errors.New("statement contains invalid verb")
  }
  if acceptedProto[strings.ToLower(splitArr[1])] {
    fmt.Println("protocol validity check: PASSED")
  } else {
    fmt.Println("protocol validity check: FAILED")
    return -1, errors.New("statement contains invalid protocol")
  }
  fmt.Println("isvalidstatement: PASSED")
  return 0, nil
}

// Build the actual iptables-compatible ruleset.
// At this point, the human input should have been validated.
func rulesetStringBuilder(humArgPtr *string)(string){
  splitArr := strings.Split(*humArgPtr, " ")
  var buf bytes.Buffer

  protoPorts := map[string]string {
    "ssh": "22",
    "http": "80",
    "https":  "443",
  }

  if strings.Compare(strings.ToLower(splitArr[1]), "icmp") != 0{
      buf.WriteString(fmt.Sprintf("-A INPUT -p tcp -m --dport %s -j %s", protoPorts[strings.ToLower(splitArr[1])], strings.ToUpper(splitArr[0])))
  } else {
    // Handling ICMP
    buf.WriteString(fmt.Sprintf("-A INPUT -p icmp -m icmp --icmp-type 0 -j %s", strings.ToUpper(splitArr[0])))
    buf.WriteString(fmt.Sprintf("\n-A INPUT -p icmp -m icmp --icmp-type 3 -j %s", strings.ToUpper(splitArr[0])))
    buf.WriteString(fmt.Sprintf("\n-A INPUT -p icmp -m icmp --icmp-type 11 -j %s", strings.ToUpper(splitArr[0])))
  }
  return buf.String()
}

func main() {
  // Defines options for the human
  var humArgPtr = flag.String("i", "null", "Accepts human definition of f/w rule.")
  flag.Parse()

  parseStdIn(humArgPtr)
  fmt.Println(rulesetStringBuilder(humArgPtr))
}