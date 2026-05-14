package main

import (
	"strings"
)

func decryptCaesar(text string, shift int) string {
	var builder strings.Builder
	for _, char := range text {
		if char < 'A' || char > 'Z' {
			builder.WriteRune(char)
			continue
		}

		offset := int(char - 'A')
		decoded := (offset - shift + len(alphabet)) % len(alphabet)
		builder.WriteByte(alphabet[decoded])
	}

	return builder.String()
}
