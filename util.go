package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"strings"
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

func CalcVariant() (varHex string) {
	trickNum = Btoi(trickTriggers[0]) | Btoi(trickTriggers[1])<<1 | Btoi(trickTriggers[2])<<2

	var (
		a = shadeNum << 3
		b = intensityNum << 6
		c = trickNum << 8
	)

	varHex = fmt.Sprintf("%03X", int64(effectType+a+b+c))

	return varHex
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CopyValue(v string) {
	v = strings.TrimPrefix(v, "0x")
	clipboard.Write(clipboard.FmtText, []byte(v))
	//if err := clipboard.WriteAll(v); err != nil {
	//	fmt.Printf("Error occurred in copying text: %v\n", err)
	//}
}
