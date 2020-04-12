package p1410

import "bytes"

func entityParser(text string) string {
	mem := make(map[string]byte)

	mem["&quot;"] = '"'
	mem["&apos;"] = '\''
	mem["&amp;"] = '&'
	mem["&gt;"] = '>'
	mem["&lt;"] = '<'
	mem["&frasl;"] = '/'

	var buf bytes.Buffer

	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			var j int
			for j < 7 && i+j < len(text) && text[i+j] != ';' {
				j++
			}
			if i+j == len(text) || j == 7 {
				// no replace
				buf.WriteByte(text[i])
				continue
			}
			// text[i+j] == ;
			if v, found := mem[text[i:i+j+1]]; !found {
				buf.WriteByte(text[i])
			} else {
				buf.WriteByte(v)
				i += j
			}
		} else {
			buf.WriteByte(text[i])
		}
	}

	return buf.String()
}
