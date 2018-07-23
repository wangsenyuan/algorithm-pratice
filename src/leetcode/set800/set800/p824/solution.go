package p824

import (
	"strings"
	"bytes"
)

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")

	a := "a"
	for i, word := range words {
		var buf bytes.Buffer
		if isVowel(word[0]) {
			buf.WriteString(word)

		} else {
			buf.WriteString(word[1:])
			buf.WriteByte(word[0])
		}
		buf.WriteString("ma")

		buf.WriteString(a)
		a += "a"
		words[i] = buf.String()
	}

	return strings.Join(words, " ")
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
