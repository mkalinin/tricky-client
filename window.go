package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

type TrickyWindow struct {
  wnd   *walk.MainWindow
  edit  *walk.TextEdit
  btn   *walk.PushButton
}

func NewTrickyWindow() (*TrickyWindow, error) {
  window := new(TrickyWindow)

  err := MainWindow{
    AssignTo: &window.wnd,
    Title:   "Tricky Client",
    MinSize: Size{600, 600},
    Layout:  VBox{},
    Children: []Widget{
      TextEdit{AssignTo: &window.edit},
      PushButton{Text: "Start", AssignTo: &window.btn},
    },
  }.Create()

  return window, err
}

func (w TrickyWindow) Run() {
  w.wnd.Run()
}

func (w TrickyWindow) AppendText(text string) {
  w.edit.SetText(w.edit.Text() + text)
}
