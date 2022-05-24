package main

import (
	"fmt"
)

func CalcFlag() (flagHex string) {
	trickNum = Btoi(trickTriggers[0]) | Btoi(trickTriggers[1])<<1 | Btoi(trickTriggers[2])<<2

	var (
		a = collisionType | effectType<<5
		b = shadeNum << 8
		c = intensityNum << 11
		d = trickNum << 13
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
