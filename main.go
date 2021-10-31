package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

var (
	colNum       int
	effNum       int
	shdNum       int
	intNum       int
	trrNum       int
	generateMenu = fyne.NewMainMenu(menu)
	flagHex      string
)

func main() {
	a := app.New()
	w := a.NewWindow("Collision Type Calculator")
	w.SetMainMenu(generateMenu)

	flagEntry := widget.NewEntry()

	colTypeSel := widget.NewSelect(types, func(s string) {
		for i := range types {
			if types[i] == s {
				colNum = i
			}
		}
		flagEntry.SetText(calcFlag())
	})
	effSel := widget.NewSelect(eff, func(s string) {
		for i := range eff {
			if eff[i] == s {
				effNum = i
			}
		}
		flagEntry.SetText(calcFlag())
	})
	shadeSel := widget.NewSelect(eff, func(s string) {
		for i := range eff {
			if eff[i] == s {
				shdNum = i
			}
		}
		flagEntry.SetText(calcFlag())
	})
	intensSel := widget.NewSelect(inte, func(s string) {
		for i := range inte {
			if inte[i] == s {
				intNum = i
			}
		}
		flagEntry.SetText(calcFlag())
	})

	w.SetContent(container.NewVBox(
		container.NewHBox(
			container.NewVBox(
				widget.NewLabel("Collision Type"),
				colTypeSel,
			),
			container.NewVBox(
				widget.NewLabel("Effect"),
				effSel,
			),
			container.NewVBox(
				widget.NewLabel("Shadow"),
				shadeSel,
			),
			container.NewVBox(
				widget.NewLabel("Intensity"),
				intensSel,
			),
			container.NewVBox(
				widget.NewLabel("Tricks"),
				widget.NewCheck(trr[0], func(b bool) {
					trrArray[0] = b
					flagEntry.SetText(calcFlag())
				}),
				widget.NewCheck(trr[1], func(b bool) {
					trrArray[1] = b
					flagEntry.SetText(calcFlag())
				}),
				widget.NewCheck(trr[2], func(b bool) {
					trrArray[2] = b
					flagEntry.SetText(calcFlag())
				}),
			)),
		container.NewVBox(flagEntry),
	))

	// Initialize option defaults
	colTypeSel.SetSelectedIndex(0)
	effSel.SetSelectedIndex(0)
	shadeSel.SetSelectedIndex(0)
	intensSel.SetSelectedIndex(0)

	// Show app
	w.ShowAndRun()
}

//func generateMenu() *fyne.MainMenu {
//	return fyne.NewMainMenu(menu)
//}

func calcFlag() string {
	trrNum = 0

	for i, b := range trrArray {
		if i == 0 && b == true {
			trrNum += 2
		} else if i == 1 && b == true {
			trrNum += 4
		} else if i == 2 && b == true {
			trrNum += 8
		}
	}

	println(trrNum)

	var (
		a = effNum * 32
		b = shdNum * 256
		c = intNum * 2048
		d = trrNum * 4096
	)

	e := strings.ToUpper(strconv.FormatInt(int64(colNum+a+b+c+d), 16))

	if len(e) <= 1 {
		flagHex = "000" + e
	} else if len(e) <= 2 {
		flagHex = "00" + e
	} else if len(e) <= 3 {
		flagHex = "0" + e
	} else {
		flagHex = e
	}

	println(flagHex)

	return flagHex
}
