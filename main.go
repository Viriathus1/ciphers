package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var alphabet = [26]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func NewInputScreen(moveToMainScreen func()) *fyne.Container {
	title := widget.NewLabel("Simple Substitution Cipher")
	title.Alignment = fyne.TextAlignCenter

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter ciphertext...")
	input.MultiLine = true

	saveButton := widget.NewButton("Save", moveToMainScreen)

	return container.NewBorder(title, saveButton, nil, nil, input)
}

func NewMainScreen() *fyne.Container {
	title := widget.NewLabel("Simple Substitution Cipher")
	title.Alignment = fyne.TextAlignCenter

	alphabetVStack := container.NewVBox()
	for _, letter := range alphabet {
		alphabetVStack.Add(container.NewBorder(nil, nil, widget.NewLabel(letter),
			container.NewHBox(widget.NewEntry(), container.NewGridWrap(fyne.Size{
				Width:  10,
				Height: 0,
			})), nil))
	}
	alphabetVScroll := container.NewVScroll(alphabetVStack)

	text := canvas.NewText("hello, world!", color.Black)
	text.Alignment = fyne.TextAlignCenter

	return container.NewBorder(title, nil, alphabetVScroll, nil, text)
}

func main() {
	a := app.New()
	w := a.NewWindow("Ciphers")

	inputScreen := NewInputScreen(func() {
		w.SetContent(NewMainScreen())
	})
	w.SetContent(inputScreen)
	w.Resize(fyne.NewSize(750, 500))
	w.ShowAndRun()
}
