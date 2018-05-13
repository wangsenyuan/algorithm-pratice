package p833

import "bytes"

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	pp := make(map[int]int)
	for i := 0; i < len(indexes); i++ {
		pp[indexes[i]] = i
	}
	var buf bytes.Buffer

	var index int

	for index < len(S) {
		if i, found := pp[index]; found {
			src := sources[i]
			tgt := targets[i]
			if index+len(src) <= len(S) && src == S[index:index+len(src)] {
				buf.WriteString(tgt)
				index += len(src)
			} else {
				buf.WriteByte(S[index])
				index++
			}
		} else {
			buf.WriteByte(S[index])
			index++
		}
	}

	return buf.String()
}
