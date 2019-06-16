package p1078

import (
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	ss := strings.Split(text, " ")

	var res []string

	for i := 0; i+2 < len(ss); i++ {
		if ss[i] == first && ss[i+1] == second {
			res = append(res, ss[i+2])
		}
	}

	return res
}
