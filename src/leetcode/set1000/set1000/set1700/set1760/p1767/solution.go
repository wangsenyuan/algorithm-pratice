package p1767

import "bytes"

func mergeAlternately(word1 string, word2 string) string {
	var buf bytes.Buffer

	var i, j int

	for i < len(word1) && j < len(word2) {
		buf.WriteByte(word1[i])
		i++
		buf.WriteByte(word2[j])
		j++
	}
	if i < len(word1) {
		buf.WriteString(word1[i:])
	}
	if j < len(word2) {
		buf.WriteString(word2[j:])
	}
	return buf.String()
}
