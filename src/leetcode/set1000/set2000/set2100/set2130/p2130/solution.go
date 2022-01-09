package p2130

import "bytes"

func capitalizeTitle(title string) string {
	var buf bytes.Buffer
	for i := 0; i < len(title); {
		if title[i] != ' ' {
			j := i
			for i < len(title) && title[i] != ' ' {
				i++
			}
			if i-j <= 2 {
				for j < i {
					buf.WriteByte(lower(title[j]))
					j++
				}
			} else {
				buf.WriteByte(upper(title[j]))
				j++
				for j < i {
					buf.WriteByte(lower(title[j]))
					j++
				}
			}
		} else {
			buf.WriteByte(title[i])
			i++
		}
	}

	return buf.String()
}

func lower(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a
	}
	return byte(a + 'a' - 'A')
}

func upper(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a
	}
	return byte(a + 'A' - 'a')
}
