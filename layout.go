package main

import (
	"fmt"
	f "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func mainWindow(a f.App) (w f.Window) {
	var (
		collisionSelect *widget.Select
		effectSelect    *widget.Select
		shadeSelect     *widget.Select
		intensitySelect *widget.Select
	)

	w = a.NewWindow("Simple KCL Calculator")

	collisionHexLabel := widget.NewLabel(fmt.Sprintf("0x%02X", collisionType))
	effectHexLabel := widget.NewLabel(fmt.Sprintf("0x%02X", effectType))

	flagEntry := widget.NewEntry()
	flagEntry.Disable()

	formLayout := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Result", Widget: flagEntry},
		},
		OnCancel: func() {
			if err := clipboard.WriteAll(flagEntry.Text); err != nil {
				fmt.Printf("Error occurred in copying text: %v\n", err)
			}
		},
		CancelText: "Copy",
	}

	collisionSelect = widget.NewSelect(collisionNames, func(s string) {
		for i, s2 := range collisionNames {
			if s2 == s {
				collisionType = i
				break
			}
		}

		collisionHexLabel.SetText(fmt.Sprintf("0x%02X", collisionType))
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
		effectHexLabel.SetText(fmt.Sprintf("0x%02X", effectType))
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
				collisionHexLabel,
			),
			container.NewVBox(
				widget.NewLabel("Effect"),
				effectSelect,
				effectHexLabel,
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
		formLayout,
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
