package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Ciphers")
	w.Resize(fyne.NewSize(750, 400))
	w.SetContent(container.NewCenter(widget.NewLabel("Choose a cipher to begin.")))
	w.Show()
	ShowCipherSelection(w)
	a.Run()
}
