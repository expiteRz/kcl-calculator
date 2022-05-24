package main

var (
	shadeOptions     = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	intensityOptions = []string{"0", "1", "2", "3"}
	trickOptions     = []string{"Trick", "No Drive", "Wall"}
	shadeNum         int
	intensityNum     int
	trickNum         int

	/// Since v0.2 ///

	collisionType  = 0
	effectType     = 0
	trickTriggers  = []bool{false, false, false}
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
		"Slippery Road 2":          {"Ice", "Mud", "Water 1", "Water 2", "Water 3", "Water 4", "Normal? 1", "Normal? 2"},
		"Boost":                    {"Default", "Used in b15", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6"},
		"Boost Ramp":               {"2 flips", "1 flip", "No flips 1", "No flips 2", "No flips 3", "No flips 4", "No flips 5", "No flips 6"},
		"Jump Pad":                 {"Stage 2 used in t72", "Stage 3 used in t53", "Stage 1 used in t62", "Stage 4 used in t13", "Mushroom", "Stage 4 used in b15", "Stage 2 used in t52 and b13", "Stage 4"},
		"Item Road":                {"Unknown 1", "Unknown 2", "Used on metals", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7"},
		"Solid Fall":               {"Sand", "Sand/Underwater", "Unknown 1", "Ice", "Dirt", "Grass", "Wood", "Unknown 2"},
		"Moving Water":             {"Water 1", "Water 2", "Water 3", "Water 4", "Asphalt 1", "Asphalt 2", "Road 1", "Road 2"},
		"Wall":                     {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Invisible Wall":           {"No spark", "Spark w/ voice 1", "Spark w/ voice 2", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5"},
		"Item Wall":                {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Wall 2":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Fall Boundary":            {"Air", "Water", "Lava", "Icy Water", "Lava", "Burning Air", "Quicksand", "Short"},
		"Cannon Activator":         {"P0", "P1", "P2", "P3", "P4", "P5", "P6", "P7"},
		"Force Recalculation":      {"N1", "N2", "N3", "N4", "N5", "N6", "N7", "N8"},
		"Half-pipe Ramp":           {"Default", "Boost applied", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6"},
		"Wall 3":                   {"Normal", "Rock", "Metal", "Wood", "Ice", "Bush", "Rope", "Rubber"},
		"Moving Road 1":            {"West 1", "East 1", "East 2", "West 2", "Rotate 1", "Rotate 2", "Unknown 1", "Unknown 2"},
		"Gravity Road":             {"Wood", "Gravel", "Carpet", "Dirt", "Sand", "Normal", "Normal", "Mud"},
		"Road 2":                   {"Normal 1", "Carpet", "Grass", "Normal 2", "Grass", "Glass", "Dirt", "Normal 3"},
		"Sound Trigger":            {"0", "1", "2", "3", "4", "5", "6", "7"},
		"Unknown":                  {"Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"},
		"Effect Trigger":           {"BRSTM Reset", "Shadow", "Water splash", "Stargate", "Halfpipe Cancel", "Coin", "Smoke", "Unknown"},
		"Item State Modifier":      {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Half-pipe Invisible Wall": {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
		"Moving Road 2":            {"Carpet", "Normal", "Normal", "Glass", "Carpet", "No sound", "Sand", "Dirt"},
		"Special Walls":            {"Cacti", "Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7"},
		"Wall 4":                   {"Unknown 1", "Unknown 2", "Unknown 3", "Unknown 4", "Unknown 5", "Unknown 6", "Unknown 7", "Unknown 8"},
	}
)
