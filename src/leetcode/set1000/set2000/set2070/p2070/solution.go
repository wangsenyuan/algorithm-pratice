package p2070

import "bytes"

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows

	var buf bytes.Buffer

	var p int
	for p < cols-rows {
		q := p
		for i := 0; i < rows; i++ {
			buf.WriteByte(encodedText[q])
			q += cols + 1
		}
		p++
	}
	for p < cols {
		q := p
		for i := p; i-p < rows && i < cols; i++ {
			buf.WriteByte(encodedText[q])
			q += cols + 1
		}
		p++
	}

	m := buf.Len()
	bs := buf.Bytes()

	for m > 0 && bs[m-1] == ' ' {
		m--
	}

	return string(bs[:m])
}
