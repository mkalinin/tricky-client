package main

import (
  "fmt"
  "log"
  "os"
)

import (
  "github.com/mattn/go-gtk/gdk"
  "github.com/mattn/go-gtk/glib"
  "github.com/mattn/go-gtk/gtk"
)

func dataLoop(buffer *gtk.TextBuffer, iter *gtk.TextIter) {
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
    gdk.ThreadsEnter()
    str := string(data)
    buffer.Insert(iter, str)
    fmt.Print(str)
    gdk.ThreadsLeave()
  }
}

func main() {
  glib.ThreadInit(nil)
  gdk.ThreadsInit()
  gdk.ThreadsEnter()
  gtk.Init(nil)

  window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
  window.SetPosition(gtk.WIN_POS_CENTER)
  window.SetTitle("Tricky Client")
  window.Connect("destroy", gtk.MainQuit)

  textview := gtk.NewTextView()
  textview.SetEditable(true)
  textview.SetCursorVisible(true)
  textview.SetSizeRequest(600, 600)
  textview.SetWrapMode(gtk.WRAP_WORD)
  
  var iter gtk.TextIter
  buffer := textview.GetBuffer()
  buffer.GetStartIter(&iter)

  window.Add(textview)
  window.SetSizeRequest(600, 600)
  window.ShowAll()

  go dataLoop(buffer, &iter)

  gtk.Main()
}
