package main

import (
  "fmt"
  "log"
  "os"
)

func dataLoop(wnd *TrickyWindow) {
  args := os.Args[1:]
  host := args[0]
  port := args[1]
  key  := args[2]

  // reader
  reader, err := NewReader(host, port)
  if err != nil {
    log.Fatal(err)
  }

  defer reader.Close()

  // decryptor
  decryptor, err := NewDecryptor(key)
  if err != nil {
    log.Fatal(err)
  }

  // read loop
  for {
    msg  := reader.Read()
    data := decryptor.Decrypt(msg)
    str  := string(data)
    
    wnd.AppendText(str)
  }
}

func main() {
  if len(os.Args) < 4 {
    fmt.Println("example of usage:")
    fmt.Println("tricky_client host port key")
    return
  }

  // creating window
  wnd, err := NewTrickyWindow()
  if err != nil {
    log.Fatal(err)
  }

  // binding start button to data loop
  wnd.btn.Clicked().Attach(func() {
    wnd.btn.SetEnabled(false)
    go dataLoop(wnd)
  })

  wnd.Run()
}
