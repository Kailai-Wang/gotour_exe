package main

import (
  "fmt"
  "strings"
  "strconv"
)

type IPAddr [4]byte

func (a IPAddr) String() string {
  var ss []string
  for _,v := range a {
    // todo: is there a better way?
    ss = append(ss, strconv.Itoa(int(v)))
  }
  return strings.Join(ss, ".")
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}

