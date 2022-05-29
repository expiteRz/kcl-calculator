package main

import "fmt"

var (
	shadeOptions     = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	intensityOptions = []string{"0", "1", "2", "3"}
	trickOptions     = []string{"Trick", "No Drive", "Wall"}
	shadeNum         int
	intensityNum     int
	trickNum         int

	/// Since v0.2 ///

	collisionType         = 0
	collisionHexFormatted = fmt.Sprintf("0x%02X", collisionType)
	effectType            = 0
	effectHexFormatted    = fmt.Sprintf("0x%03X", effectType)
	trickTriggers         = []bool{false, false, false}
	collisionNames        = []string{
		"Road (0x00)",
		"Slippery Road 1 (0x01)",
		"Weak Off-road (0x02)",
		"Off-road (0x03)",
		"Heavy Off-road (0x04)",
		"Slippery Road 2 (0x05)",
		"Boost (0x06)",
		"Boost Ramp (0x07)",
		"Jump Pad (0x08)",
		"Item Road (0x09)",
		"Solid Fall (0x0A)",
		"Moving Water (0x0B)",
		"Wall (0x0C)",
		"Invisible Wall (0x0D)",
		"Item Wall (0x0E)",
		"Wall 2 (0x0F)",
		"Fall Boundary (0x10)",
		"Cannon Activator (0x11)",
		"Force Recalculation (0x12)",
		"Half-pipe Ramp (0x13)",
		"Wall 3 (0x14)",
		"Moving Road 1 (0x15)",
		"Gravity Road (0x16)",
		"Road 2 (0x17)",
		"Sound Trigger (0x18)",
		"Weak Wall (0x19)",
		"Effect Trigger (0x1A)",
		"Item State Modifier (0x1B)",
		"Half-pipe Invisible Wall (0x1C)",
		"Moving Road 2 (0x1D)",
		"Special Walls (0x1E)",
		"Wall 4 (0x1F)",
	}
	effectTypes = map[string][]string{
		"Road (0x00)":                     {"Normal", "Dirt w/ GFX", "Dirt w/o GFX", "Smooth", "Wood", "Snow w/ GFX", "Metal", "Normal?"},
		"Slippery Road 1 (0x01)":          {"White Sand", "Dirt 1", "Water", "Snow", "Grass", "Yellow Sand", "Sand", "Dirt 2"},
		"Weak Off-road (0x02)":            {"Orange Sand", "Dirt", "Water", "Grass", "Sand", "Carpet", "Gravel 1", "Gravel 2"},
		"Off-road (0x03)":                 {"Sand 1", "Dirt", "Mud", "Water", "Grass", "Sand 2", "Gravel", "Carpet"},
		"Heavy Off-road (0x04)":           {"Sand 1", "Dirt", "Mud", "Flower", "Grass", "Snow", "Sand 2", "Dirt"},
		"Slippery Road 2 (0x05)":          {"Ice", "Mud", "Water 1", "Water 2", "Water 3", "Water 4", "Normal? 1", "Normal? 2"},
		"Boost (0x06)":                    {"Default", "Used in b15", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6"},
		"Boost Ramp (0x07)":               {"2 flips", "1 flip", "No flips 1", "No flips 2", "No flips 3", "No flips 4", "No flips 5", "No flips 6"},
		"Jump Pad (0x08)":                 {"Stage 2 used in t72", "Stage 3 used in t53", "Stage 1 used in t62", "Stage 4 used in t13", "Mushroom", "Stage 4 used in b15", "Stage 2 used in t52 and b13", "Stage 4"},
		"Item Road (0x09)":                {"Unknown 1", "Unknown 2", "Used on metals", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7"},
		"Solid Fall (0x0A)":               {"Sand", "Sand/Underwater", "Unknown 1", "Ice", "Dirt", "Grass", "Wood", "Unknown 2"},
		"Moving Water (0x0B)":             {"Water 1", "Water 2", "Water 3", "Water 4", "Asphalt 1", "Asphalt 2", "Road 1", "Road 2"},
		"Wall (0x0C)":                     {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Invisible Wall (0x0D)":           {"No spark", "Spark w/ voice 1", "Spark w/ voice 2", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5"},
		"Item Wall (0x0E)":                {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Wall 2 (0x0F)":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Fall Boundary (0x10)":            {"Air", "Water", "Lava", "Icy Water", "Lava", "Burning Air", "Quicksand", "Short"},
		"Cannon Activator (0x11)":         {"P0", "P1", "P2", "P3", "P4", "P5", "P6", "P7"},
		"Force Recalculation (0x12)":      {"N1", "N2", "N3", "N4", "N5", "N6", "N7", "N8"},
		"Half-pipe Ramp (0x13)":           {"Default", "Boost applied", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6"},
		"Wall 3 (0x14)":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Moving Road 1 (0x15)":            {"West 1", "East 1", "East 2", "West 2", "Rotate 1", "Rotate 2", "Unknown 1", "Unknown 2"},
		"Gravity Road (0x16)":             {"Wood", "Gravel", "Carpet", "Dirt", "Sand", "Rainbow Road", "Normal", "Mud"},
		"Road 2 (0x17)":                   {"Normal 1", "Carpet", "Grass", "Normal 2", "Grass", "Glass", "Dirt", "Rainbow Road"},
		"Sound Trigger (0x18)":            {"0", "1", "2", "3", "4", "5", "6", "7"},
		"Weak Wall (0x19)":                {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Effect Trigger (0x1A)":           {"BRSTM Reset", "Shadow", "Water splash", "Stargate", "Halfpipe Cancel", "Coin", "Smoke", "Unknown"},
		"Item State Modifier (0x1B)":      {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Half-pipe Invisible Wall (0x1C)": {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Moving Road 2 (0x1D)":            {"Carpet 1", "Normal 1", "Normal 2", "Glass", "Carpet 2", "Star Crash Impact", "Sand", "Dirt"},
		"Special Walls (0x1E)":            {"Cacti", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7"},
		"Wall 4 (0x1F)":                   {"Fast Bump", "Unused 1", "Unused 2", "Unused 3", "Unused 4", "Unused 5", "Unused 6", "Unused 7"},
	}
)
