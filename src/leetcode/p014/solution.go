package main

import (
	"fmt"
	"sort"
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

	for _, str := range strs {
		prefix = appendStrPrefix(prefix, str)
	}

	sort.Strings(prefix)

	return findCommonPrefix(prefix, len(strs))
}

func findCommonPrefix(pres []string, n int) string {
	i, j := 0, len(pres)-1

	for i <= j {
		m := (i + j) / 2

		c := countOfPreAt(pres, m)
		if c == n {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return pres[i-1]
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

func appendStrPrefix(pre []string, str string) []string {
	for i := 0; i <= len(str); i++ {
		pre = appendMore(pre, str[:i])
	}
	return pre
}

func appendMore(strs []string, str string) []string {
	if len(strs)+1 == cap(strs) {
		tmp := make([]string, len(strs), 2*cap(strs))
		copy(tmp, strs)
		strs = tmp
	}
	return append(strs, str)
}
