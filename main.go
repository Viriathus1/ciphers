package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Ciphers")

	inputScreen := NewInputScreen(func() {
		w.SetContent(NewMainScreen())
	})
	w.SetContent(inputScreen)
	w.Resize(fyne.NewSize(750, 300))
	w.ShowAndRun()
}
