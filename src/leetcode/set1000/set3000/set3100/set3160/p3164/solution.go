package p3164

import (
	"bytes"
	"fmt"
)

func compressedString(word string) string {
	var buf bytes.Buffer

	for i := 0; i < len(word); {
		j := i
		for i < len(word) && i < j+9 && word[i] == word[j] {
			i++
		}
		buf.WriteString(fmt.Sprintf("%d%c", i-j, word[j]))
	}

	return buf.String()
}
