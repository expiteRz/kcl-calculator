package main

import (
	"fyne.io/fyne/v2"
)

var (
	types, _ = GetIndexes(DB{0, 0})
	shd      = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	inte     = []string{"0", "1", "2", "3"}
	menu     = fyne.NewMenu("File")
	trr      = []string{"Trick", "No Drive", "Wall"}
	trrArray = []bool{false, false, false}
)

type DataMode int

const (
	colType = iota // collision
	effType        // effect
)
