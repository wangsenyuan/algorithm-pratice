package p820

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	n := len(words)
	rw := make([]string, n)

	for i := 0; i < n; i++ {
		rw[i] = reverse(words[i])
	}
	sort.Strings(rw)

	var ans int

	for i := 1; i <= n; i++ {
		if i < n && strings.HasPrefix(rw[i], rw[i-1]) {
			continue
		}
		ans += len(rw[i-1]) + 1
	}

	return ans
}

func reverse(s string) string {
	n := len(s)
	r := []byte(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
