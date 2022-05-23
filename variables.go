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

var (
	collisionType  = 0
	effectType     = 0
	collisionNames = []string{
		"Road",
		"Slippery Road 1",
		"Weak Off-road",
		"Off-road",
		"Heavy Off-road",
		"Slippery Road 2",
		"Boost",
		"Boost Ramp",
		"Jump Pad",
		"Item Road",
		"Solid Fall",
		"Moving Water",
		"Wall",
		"Invisible Wall",
		"Item Wall",
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
	effectTypes = map[string][]string{
		"Road":                     {"Normal", "Dirt w/ GFX", "Dirt w/o GFX", "Smooth", "Wood", "Snow w/ GFX", "Metal", "Normal?"},
		"Slippery Road 1":          {"White Sand", "Dirt", "Water", "Snow", "Grass", "Yellow Sand", "Sand", "Dirt"},
		"Weak Off-road":            {"Orange Sand", "Dirt", "Water", "Grass", "Sand", "Carpet", "Gravel", "Gravel"},
		"Off-road":                 {"Sand", "Dirt", "Mud", "Water", "Grass", "Sand", "Gravel", "Carpet"},
		"Heavy Off-road":           {"Sand", "Dirt", "Mud", "Flower", "Grass", "Snow", "Sand", "Dirt"},
		"Slippery Road 2":          {"Ice", "Mud", "Water", "Water", "Water", "Water", "Normal?", "Normal?"},
		"Boost":                    {"Default", "Used in b15", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Boost Ramp":               {"2 flips", "1 flip", "No flips", "No flips", "No flips", "No flips", "No flips", "No flips"},
		"Jump Pad":                 {"Stage 2 used in t72", "Stage 3 used in t53", "Stage 1 used in t62", "Stage 4 used in t13", "Mushroom", "Stage 4 used in b15", "Stage 2 used in t52 and b13", "Stage 4"},
		"Item Road":                {"Unknown", "Unknown", "Used on metals", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Solid Fall":               {"Sand", "Sand/Underwater", "Unknown", "Ice", "Dirt", "Grass", "Wood", "Unknown"},
		"Moving Water":             {"Water", "Water", "Water", "Water", "Asphalt", "Asphalt", "Road", "Road"},
		"Wall":                     {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Invisible Wall":           {"No spark", "Spark w/ voice", "Spark w/ voice", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Item Wall":                {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Wall 2":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Fall Boundary":            {"Air", "Water", "Lava", "Icy Water", "Lava", "Burning Air", "Quicksand", "Short"},
		"Cannon Activator":         {"P0", "P1", "P2", "P3", "P4", "P5", "P6", "P7"},
		"Force Recalculation":      {"N", "N", "N", "N", "N", "N", "N", "N"},
		"Half-pipe Ramp":           {"Default", "Boost applied", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Wall 3":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Moving Road 1":            {"West", "East", "East", "West", "Rotate", "Rotate", "Unknown", "Unknown"},
		"Gravity Road":             {"Wood", "Gravel", "Carpet", "Dirt", "Sand", "Normal", "Normal", "Mud"},
		"Road 2":                   {"Normal", "Carpet", "Grass", "Normal", "Grass", "Glass", "Dirt", "Normal"},
		"Sound Trigger":            {"0", "1", "2", "3", "4", "5", "6", "7"},
		"Unknown":                  {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Effect Trigger":           {"BRSTM Reset", "Shadow", "Water splash", "Stargate", "Halfpipe Cancel", "Coin", "Smoke", "Unknown"},
		"Item State Modifier":      {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Half-pipe Invisible Wall": {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Moving Road 2":            {"Carpet", "Normal", "Normal", "Glass", "Carpet", "No sound", "Sand", "Dirt"},
		"Special Walls":            {"Cacti", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Wall 4":                   {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
	}
)
