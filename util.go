package main

import (
	"fmt"
)

func CalcFlag() (flagHex string) {
	trrNum = Btoi(trrArray[0]) | Btoi(trrArray[1])<<1 | Btoi(trrArray[2])<<2

	var (
		a = collisionType | effectType<<5
		b = shdNum << 8
		c = intNum << 11
		d = trrNum << 13
	)

	flagHex = fmt.Sprintf("%04X", int64(a+b+c+d))

	return flagHex
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
