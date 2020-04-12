package p1408

import (
	"strings"
)

func stringMatching(words []string) []string {

	var res []string

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) && len(words[j]) > len(words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}

	return res
}
