package p1309

import "bytes"

func freqAlphabets(s string) string {
	ii := make(map[string]byte)

	ii["1"] = 'a'
	ii["2"] = 'b'
	ii["3"] = 'c'
	ii["4"] = 'd'
	ii["5"] = 'e'
	ii["6"] = 'f'
	ii["7"] = 'g'
	ii["8"] = 'h'
	ii["9"] = 'i'
	ii["10#"] = 'j'
	ii["11#"] = 'k'
	ii["12#"] = 'l'
	ii["13#"] = 'm'
	ii["14#"] = 'n'
	ii["15#"] = 'o'
	ii["16#"] = 'p'
	ii["17#"] = 'q'
	ii["18#"] = 'r'
	ii["19#"] = 's'
	ii["20#"] = 't'
	ii["21#"] = 'u'
	ii["22#"] = 'v'
	ii["23#"] = 'w'
	ii["24#"] = 'x'
	ii["25#"] = 'y'
	ii["26#"] = 'z'

	var buf bytes.Buffer

	var i int

	for i < len(s) {
		x := s[i]
		if i >= len(s)-2 || (x >= '3' && x <= '9') {
			buf.WriteByte(ii[s[i:i+1]])
			i++
			continue
		}
		z := s[i+2]
		if z != '#' {
			buf.WriteByte(ii[s[i:i+1]])
			i++
			continue
		}
		buf.WriteByte(ii[s[i:i+3]])
		i += 3
	}

	return buf.String()
}
