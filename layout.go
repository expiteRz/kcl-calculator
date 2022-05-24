package main

import (
	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainWindow(a f.App) (w f.Window) {
	var (
		collisionSelect *widget.Select
		effectSelect    *widget.Select
		shadeSelect     *widget.Select
		intensitySelect *widget.Select
	)

	w = a.NewWindow("Simple KCL Calculator")

	flagEntry := widget.NewEntry()
	flagEntry.Disable()

	collisionSelect = widget.NewSelect(collisionNames, func(s string) {
		for i, s2 := range collisionNames {
			if s2 == s {
				collisionType = i
				break
			}
		}

		selectedEffect := effectSelect.SelectedIndex()
		effectSelect.Options = effectTypes[collisionNames[collisionType]]
		effectSelect.Refresh()
		effectSelect.SetSelectedIndex(selectedEffect)
		flagEntry.SetText(CalcFlag())
	})

	effectSelect = widget.NewSelect(effectTypes[collisionNames[collisionType]], func(s string) {
		for i, s2 := range effectTypes[collisionNames[collisionType]] {
			if s2 == s {
				effectType = i
			}
		}
		flagEntry.SetText(CalcFlag())
	})

	shadeSelect = widget.NewSelect(shadeOptions, func(s string) {
		for i := range shadeOptions {
			if shadeOptions[i] == s {
				shadeNum = i
			}
		}
		flagEntry.SetText(CalcFlag())
	})

	intensitySelect = widget.NewSelect(intensityOptions, func(s string) {
		for i := range intensityOptions {
			if intensityOptions[i] == s {
				intensityNum = i
			}
		}
		flagEntry.SetText(CalcFlag())
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
				widget.NewCheck(trickOptions[0], func(b bool) {
					trickTriggers[0] = b
					flagEntry.SetText(CalcFlag())
				}),
				widget.NewCheck(trickOptions[1], func(b bool) {
					trickTriggers[1] = b
					flagEntry.SetText(CalcFlag())
				}),
				widget.NewCheck(trickOptions[2], func(b bool) {
					trickTriggers[2] = b
					flagEntry.SetText(CalcFlag())
				}),
			)),
		container.NewVBox(flagEntry),
	))

	collisionSelect.SetSelectedIndex(0)
	effectSelect.SetSelectedIndex(0)
	shadeSelect.SetSelectedIndex(0)
	intensitySelect.SetSelectedIndex(0)

	size := w.Canvas().Size()
	w.SetFixedSize(true)
	w.Resize(size)

	return w
}
