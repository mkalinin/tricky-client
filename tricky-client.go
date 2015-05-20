package main

import (
  "fmt"
  "log"
  "os"
)

func main() {
  args := os.Args[1:]

  host := args[0]
  port := args[1]
  key  := args[2]

  reader, err := NewReader(host, port)
  if err != nil {
    log.Fatal(err)
  }

  defer reader.Close()

  // setting up decryptor
  decryptor, err := NewDecryptor(key)
  if err != nil {
    log.Fatal(err)
    fmt.Print(err)
  }

  for {
    msg  := reader.Read()
    data := decryptor.Decrypt(msg)
    fmt.Print(string(data))
  }
}
