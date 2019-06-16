package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("%s\n", longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := make([]string, 0, len(strs))

	minLen := math.MaxInt32

	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	for _, str := range strs {
		prefix = appendStrPrefix(prefix, str, minLen)
	}

	sort.Strings(prefix)

	return findCommonPrefix(prefix, len(strs))
}

func findCommonPrefix(pres []string, n int) string {
	i, j := 0, len(pres) - 1

	for i <= j {
		m := (i + j) / 2

		c := countOfPreAt(pres, m)
		if c == n {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return pres[i - 1]
}

func countOfPreAt(pres []string, p int) int {
	s := pres[p]

	i := p
	for i >= 0 && pres[i] == s {
		i--
	}

	j := p
	for j < len(pres) && pres[j] == s {
		j++
	}

	return (j - 1) - (i + 1) + 1
}

func appendStrPrefix(pre []string, str string, minLen int) []string {
	for i := 0; i <= minLen; i++ {
		pre = appendMore(pre, str[:i])
	}
	return pre
}

func appendMore(strs []string, str string) []string {
	if len(strs) + 1 == cap(strs) {
		tmp := make([]string, len(strs), 2 * cap(strs))
		copy(tmp, strs)
		strs = tmp
	}
	return append(strs, str)
}
