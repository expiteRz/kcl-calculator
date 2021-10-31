package main

import "fyne.io/fyne/v2"

type DialogType int

const (
	info = iota
	warn
)

type Dialog struct {
	app fyne.App
}

func (d Dialog) createDialog(msg string, dialogType DialogType) {
	d.app.NewWindow("Test").Show()
}
