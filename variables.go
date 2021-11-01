package main

import (
	"fyne.io/fyne/v2"
)

var (
	types    = DB{mode: colType, effectId: 0}.GetIndexes()
	eff      []string
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
