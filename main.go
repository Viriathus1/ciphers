package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewInputScreen(moveToMainScreen func()) *fyne.Container {
	title := widget.NewLabel("Simple Substitution Cipher")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter ciphertext...")
	input.MultiLine = true

	saveButton := widget.NewButton("Save", moveToMainScreen)

	return container.NewBorder(title, saveButton, nil, nil, input)
}

func main() {
	a := app.New()
	w := a.NewWindow("Ciphers")

	inputScreen := NewInputScreen(func() {
		w.SetContent(container.NewWithoutLayout())
	})
	w.SetContent(inputScreen)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
