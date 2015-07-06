// script arguments: location key id caveat_id

package main

import (
  "os"
  "fmt"
  b64 "encoding/base64"
  "strings"

  "gopkg.in/macaroon.v1"
)


func main() {
    args := os.Args[1:]

    key := []byte(args[1])

    m, err := macaroon.New(key, args[2], args[0])

    m.AddFirstPartyCaveat(args[3])

    serialized, err := m.MarshalBinary()
    encoded := strings.TrimRight(b64.URLEncoding.EncodeToString(serialized), "=")

    if err == nil {
      fmt.Printf("%s\n%s\n%s\n", encoded, key, args[3])
    }
}
