package main

import (
	"fyne.io/fyne/v2"
)

var (
	types = []string{
		"Road",
		"Slippery Road 1",
		"Weak Off-road",
		"Off-road",
		"Heavy Off-road",
		"Slippery Road 2",
		"Boost Pad",
		"Boost Ramp",
		"Jump Pad",
		"Item Road",
		"Solid Fall",
		"Moving Water",
		"Wall",
		"Invisible Wall",
		"Item Item",
		"Wall 2",
		"Fall Boundary",
		"Cannon Activator",
		"Force Recalculation",
		"Half-pipe Ramp",
		"Wall 3",
		"Moving Road",
		"Gravity Road",
		"Road 2",
		"Sound Trigger",
		"Unknown",
		"Effect Trigger",
		"Item State Modifier",
		"Half-pipe Invisible Wall",
		"Moving Road",
		"Special Walls",
		"Wall 4",
	}
	eff      = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	inte     = []string{"0", "1", "2", "3"}
	menu     = fyne.NewMenu("File")
	trr      = []string{"Trick", "No Drive", "Wall"}
	trrArray = []bool{false, false, false}
)
