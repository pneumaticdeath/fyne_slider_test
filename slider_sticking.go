package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Slider Test")

	data := binding.NewFloat()
	slider := widget.NewSliderWithData(0.0, 10.0, data)
	slider.SetValue(5.0)
	slider.Step = 0.1
	display := widget.NewLabel("Holding")
	data.AddListener(binding.NewDataListener(func() {
		val, _ := data.Get()
		display.Text = fmt.Sprint(val)
		display.Refresh()
	}))

	content := container.New(layout.NewVBoxLayout(), slider, display)

	w.SetContent(content)
	w.Resize(fyne.NewSize(100,100))
	w.ShowAndRun()
}

