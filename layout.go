package main

import (
	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainWindow(a f.App) {
	var (
		collisionSelect *widget.Select
		effectSelect    *widget.Select
		shadeSelect     *widget.Select
		intensitySelect *widget.Select
	)

	w := a.NewWindow("Simple KCL Calculator")

	flagEntry := widget.NewEntry()
	flagEntry.Disable()

	collisionSelect = widget.NewSelect(collisionNames, func(s string) {
		for i := range collisionNames {
			collisionType = i
		}

		selectedEffect := effectSelect.SelectedIndex()
		effectSelect.Options = effectTypes[collisionNames[collisionType]]
		effectSelect.Refresh()
		effectSelect.SetSelectedIndex(selectedEffect)
	})

	effectSelect = widget.NewSelect(effectTypes[collisionNames[collisionType]], func(s string) {
		for i, s2 := range effectTypes[collisionNames[collisionType]] {
			if s2 == s {
				effNum = i
			}
		}
	})

	shadeSelect = widget.NewSelect(shd, func(s string) {
		for i := range shd {
			if shd[i] == s {
				shdNum = i
			}
		}
	})

	intensitySelect = widget.NewSelect(inte, func(s string) {
		for i := range inte {
			if inte[i] == s {
				intNum = i
			}
		}
	})

	// Assign widgets into a window
	w.SetContent(container.NewVBox(
		container.NewHBox(
			container.NewVBox(
				widget.NewLabel("Collision Type"),
				collisionSelect,
			),
			container.NewVBox(
				widget.NewLabel("Effect"),
				effectSelect,
			),
			container.NewVBox(
				widget.NewLabel("Shadow"),
				shadeSelect,
			),
			container.NewVBox(
				widget.NewLabel("Intensity"),
				intensitySelect,
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
}
