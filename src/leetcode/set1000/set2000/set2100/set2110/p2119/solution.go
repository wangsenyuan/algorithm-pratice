package p2119

import "strings"

func mostWordsFound(sentences []string) int {
	var res int

	for _, sent := range sentences {
		res = max(res, len(strings.Split(sent, " ")))
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
