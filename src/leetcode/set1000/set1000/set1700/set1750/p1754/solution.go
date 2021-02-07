package p1754

import "bytes"

func largestMerge(word1 string, word2 string) string {
	var buf bytes.Buffer

	var i, j int

	for i < len(word1) && j < len(word2) {
		if word1[i] > word2[j] {
			buf.WriteByte(word1[i])
			i++
		} else if word1[i] < word2[j] {
			buf.WriteByte(word2[j])
			j++
		} else {
			// we need to find the first different one
			ii := i
			jj := j
			for ii < len(word1) && jj < len(word2) && word1[ii] == word2[jj] {
				ii++
				jj++
			}
			if ii == len(word1) || (jj < len(word2) && word2[jj] > word1[ii]) {
				buf.WriteByte(word2[j])
				j++
			} else {
				buf.WriteByte(word1[i])
				i++
			}
		}
	}

	if i < len(word1) {
		buf.WriteString(word1[i:])
	}
	if j < len(word2) {
		buf.WriteString(word2[j:])
	}

	return buf.String()
}
