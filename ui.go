package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func ShowCipherSelection(w fyne.Window) {
	var popup *widget.PopUp

	openSubstitution := func() {
		popup.Hide()
		w.SetContent(NewInputScreen("Simple Substitution Cipher", func() {
			ShowCipherSelection(w)
		}, func(ciphertext binding.String) {
			w.SetContent(NewSubstitutionScreen(ciphertext, func() {
				ShowCipherSelection(w)
			}))
		}))
	}

	openCaesar := func() {
		popup.Hide()
		w.SetContent(NewInputScreen("Caesar Cipher", func() {
			ShowCipherSelection(w)
		}, func(ciphertext binding.String) {
			w.SetContent(NewCaesarScreen(ciphertext, func() {
				ShowCipherSelection(w)
			}))
		}))
	}

	title := widget.NewLabel("Choose a cipher")
	title.Alignment = fyne.TextAlignCenter

	content := container.NewPadded(container.NewVBox(
		title,
		widget.NewButton("Simple Substitution Cipher", openSubstitution),
		widget.NewButton("Caesar Cipher", openCaesar),
	))

	popup = widget.NewModalPopUp(content, w.Canvas())
	popup.Resize(fyne.NewSize(320, 160))
	popup.Show()
}

func NewInputScreen(titleText string, onBack func(), onSubmit func(binding.String)) fyne.CanvasObject {
	ciphertext := binding.NewString()

	title := widget.NewLabel(titleText)
	title.Alignment = fyne.TextAlignCenter

	input := widget.NewEntryWithData(ciphertext)
	input.SetPlaceHolder("Enter ciphertext...")
	input.MultiLine = true
	input.Validator = validateCiphertext
	input.AlwaysShowValidationError = true
	input.Wrapping = fyne.TextWrapBreak
	input.SetMinRowsVisible(10)

	form := &widget.Form{
		OnSubmit: func() {
			if err := input.Validate(); err == nil {
				onSubmit(ciphertext)
			}
		},
		SubmitText: "Continue",
	}
	form.Append("Ciphertext", input)

	return container.NewBorder(
		container.NewVBox(widget.NewButton("Change Cipher", onBack), title),
		nil,
		nil,
		nil,
		form,
	)
}

func NewSubstitutionScreen(ciphertext binding.String, onBack func()) fyne.CanvasObject {
	normalizedText := normalizeCiphertext(ciphertext)
	plaintext := binding.NewString()

	for _, char := range alphabet {
		keymap[char] = char
	}

	updatePreview := func() {
		text, _ := ciphertext.Get()
		plaintext.Set(decryptSubstitution(text, keymap))
	}

	updatePreview()

	title := widget.NewLabel("Simple Substitution Cipher")
	title.Alignment = fyne.TextAlignCenter

	alphabetVStack := container.NewVBox()
	for _, letter := range alphabet {
		letterField := widget.NewEntry()
		letterField.SetPlaceHolder(strings.ToLower(string(letter)))
		letterField.OnChanged = func(s string) {
			firstRune, _ := utf8.DecodeRuneInString(s)
			if firstRune == utf8.RuneError {
				keymap[letter] = letter
				updatePreview()
				return
			}

			normalizedRune := unicode.ToLower(firstRune)
			normalizedText := string(normalizedRune)
			if s != normalizedText {
				letterField.SetText(normalizedText)
				return
			}

			keymap[letter] = normalizedRune
			updatePreview()
		}
		alphabetVStack.Add(container.NewBorder(nil, nil, widget.NewLabel(string(letter)),
			container.NewHBox(letterField, container.NewGridWrap(fyne.Size{
				Width:  10,
				Height: 0,
			})), nil))
	}
	alphabetVScroll := container.NewVScroll(alphabetVStack)

	plaintext.Set(normalizedText)
	text := widget.NewLabelWithData(plaintext)
	text.Alignment = fyne.TextAlignCenter
	text.Wrapping = fyne.TextWrapBreak

	return container.NewBorder(
		container.NewVBox(widget.NewButton("Change Cipher", onBack), title),
		nil,
		alphabetVScroll,
		nil,
		text,
	)
}

func NewCaesarScreen(ciphertext binding.String, onBack func()) fyne.CanvasObject {
	normalizedText := normalizeCiphertext(ciphertext)
	plaintext := binding.NewString()

	shiftLabel := widget.NewLabel("Shift: 0")
	shiftLabel.Alignment = fyne.TextAlignCenter

	updatePreview := func(shift int) {
		plaintext.Set(decryptCaesar(normalizedText, shift))
		shiftLabel.SetText(fmt.Sprintf("Shift: %d", shift))
	}

	updatePreview(0)

	title := widget.NewLabel("Caesar Cipher")
	title.Alignment = fyne.TextAlignCenter

	slider := widget.NewSlider(0, 25)
	slider.Step = 1
	slider.OnChanged = func(value float64) {
		updatePreview(int(value))
	}

	ciphertextTitle := widget.NewLabel("Ciphertext")
	ciphertextTitle.Alignment = fyne.TextAlignCenter

	ciphertextLabel := widget.NewLabel(normalizedText)
	ciphertextLabel.Alignment = fyne.TextAlignCenter
	ciphertextLabel.Wrapping = fyne.TextWrapBreak

	plaintextTitle := widget.NewLabel("Plaintext")
	plaintextTitle.Alignment = fyne.TextAlignCenter

	plaintextLabel := widget.NewLabelWithData(plaintext)
	plaintextLabel.Alignment = fyne.TextAlignCenter
	plaintextLabel.Wrapping = fyne.TextWrapBreak

	return container.NewBorder(
		container.NewVBox(widget.NewButton("Change Cipher", onBack), title),
		nil,
		nil,
		nil,
		container.NewVBox(
			shiftLabel,
			slider,
			ciphertextTitle,
			ciphertextLabel,
			plaintextTitle,
			plaintextLabel,
		),
	)
}

func validateCiphertext(s string) error {
	if len(s) == 0 {
		return errors.New("ciphertext is too small")
	}
	if len(s) > 1000 {
		return errors.New("ciphertext is too BIG")
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 128 {
			return errors.New("invalid ascii string")
		}
	}

	return nil
}

func normalizeCiphertext(ciphertext binding.String) string {
	ciphertextContent, _ := ciphertext.Get()
	normalizedText := strings.ToUpper(ciphertextContent)
	ciphertext.Set(normalizedText)
	return normalizedText
}
