package main

import (
	"fyne.io/fyne/v2/app"
	"golang.design/x/clipboard"
)

func main() {
	a := app.New()
	if err := clipboard.Init(); err != nil {
		return
	}
	w := mainWindow(a)
	w.ShowAndRun()
}
