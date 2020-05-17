package p1450

import (
	"bytes"
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	if len(text) == 0 {
		return text
	}
	words := strings.Split(text, " ")
	ws := make([]Word, len(words))

	for i := 0; i < len(words); i++ {
		ws[i] = Word{words[i], i}
	}

	sort.Sort(Words(ws))

	var buf bytes.Buffer

	buf.WriteString(ws[0].Capital())

	for i := 1; i < len(ws); i++ {
		buf.WriteByte(' ')
		buf.WriteString(ws[i].Lower())
	}

	return buf.String()
}

type Word struct {
	word string
	idx  int
}

func (word Word) Capital() string {
	buf := []byte(word.word)

	if buf[0] >= 'a' && buf[0] <= 'z' {
		buf[0] = buf[0] - 'a' + 'A'
	}

	return string(buf)
}

func (word Word) Lower() string {
	buf := []byte(word.word)

	if buf[0] >= 'A' && buf[0] <= 'Z' {
		buf[0] = buf[0] - 'A' + 'a'
	}

	return string(buf)
}

type Words []Word

func (words Words) Len() int {
	return len(words)
}

func (words Words) Less(i, j int) bool {
	if len(words[i].word) < len(words[j].word) {
		return true
	}
	if len(words[i].word) == len(words[j].word) {
		return words[i].idx < words[j].idx
	}
	return false
}

func (words Words) Swap(i, j int) {
	words[i], words[j] = words[j], words[i]
}
