package main

import (
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var keymap = make(map[rune]rune, len(alphabet))

func decryptSubstitution(text string, keyMap map[rune]rune) string {
	var builder strings.Builder
	for _, char := range text {
		if val, ok := keyMap[char]; ok {
			builder.WriteRune(val)
			continue
		}
		builder.WriteRune(char)
	}

	return builder.String()
}
