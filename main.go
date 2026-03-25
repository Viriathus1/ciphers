package main

import (
	"strings"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
	keyMap     = make(map[rune]rune)
	ciphertext = binding.NewString()
	plaintext  = binding.NewString()
)

func NewInputScreen(moveToMainScreen func()) *fyne.Container {
	title := widget.NewLabel("Simple Substitution Cipher")
	title.Alignment = fyne.TextAlignCenter

	input := widget.NewEntryWithData(ciphertext)
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
		letterField := widget.NewEntry()
		letterField.OnChanged = func(s string) {
			firstRune, _ := utf8.DecodeRuneInString(s)
			if firstRune == utf8.RuneError {
				keyMap[letter] = letter
			} else {
				keyMap[letter] = firstRune
			}
			text, _ := ciphertext.Get()
			var builder strings.Builder
			for _, char := range text {
				// need to check for spaces, punction etc.
				builder.WriteRune(keyMap[char])
			}
			plaintext.Set(builder.String())
		}
		alphabetVStack.Add(container.NewBorder(nil, nil, widget.NewLabel(string(letter)),
			container.NewHBox(letterField, container.NewGridWrap(fyne.Size{
				Width:  10,
				Height: 0,
			})), nil))
	}
	alphabetVScroll := container.NewVScroll(alphabetVStack)

	text := widget.NewLabelWithData(plaintext)
	text.Alignment = fyne.TextAlignCenter

	return container.NewBorder(title, nil, alphabetVScroll, nil, text)
}

func main() {
	a := app.New()
	w := a.NewWindow("Ciphers")

	inputScreen := NewInputScreen(func() {
		for _, char := range alphabet {
			keyMap[char] = char
		}
		ciphertextContent, _ := ciphertext.Get()
		loweredText := strings.ToLower(ciphertextContent)
		ciphertext.Set(loweredText)
		plaintext.Set(loweredText)
		w.SetContent(NewMainScreen())
	})
	w.SetContent(inputScreen)
	w.Resize(fyne.NewSize(750, 500))
	w.ShowAndRun()
}
