package p888

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	a := diff(A, B)
	b := diff(B, A)
	return append(a, b...)
}

func diff(A, B string) []string {
	cache := make(map[string]bool)
	ss := strings.Split(B, " ")
	for _, s := range ss {
		cache[s] = true
	}
	as := strings.Split(A, " ")
	cnt := make(map[string]int)
	for _, a := range as {
		cnt[a]++
	}
	res := make([]string, 0, len(as))
	for _, a := range as {
		if cnt[a] == 1 && cache[a] == false {
			res = append(res, a)
		}
	}
	return res
}
