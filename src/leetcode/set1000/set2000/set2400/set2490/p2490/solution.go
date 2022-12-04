package p2490

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")

	n := len(words)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		m := len(words[i])
		if words[i][m-1] != words[j][0] {
			return false
		}
	}
	return true
}
