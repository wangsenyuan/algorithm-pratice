package p2788

import (
	"fmt"
	"strings"
)

func splitWordsBySeparator(words []string, separator byte) []string {
	var res []string

	sep := fmt.Sprintf("%c", separator)

	for _, word := range words {
		ss := strings.Split(word, sep)

		for _, s := range ss {
			if len(s) > 0 {
				res = append(res, s)
			}
		}
	}
	return res
}
