package main

import "strconv"

func CalcFlag() (flagHex string) {
	trrNum = Btoi(trrArray[0]) | Btoi(trrArray[1])<<1 | Btoi(trrArray[2])<<2

	var (
		a = collisionType | effectType<<5
		b = shdNum << 8
		c = intNum << 11
		d = trrNum << 13
	)

	flagFormat := strconv.FormatInt(int64(a+b+c+d), 16)

	switch len(flagFormat) {
	case 1:
		flagHex = "000" + flagFormat
	case 2:
		flagHex = "00" + flagFormat
	case 3:
		flagHex = "0" + flagFormat
	default:
		return flagFormat
	}

	return flagHex
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
