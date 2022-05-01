package p2255

import "sort"

func countPrefixes(words []string, s string) int {
	sort.Strings(words)
	var i int
	for i < len(words) && words[i][0] < s[0] {
		i++
	}

	if i == len(words) || words[i][0] > s[0] {
		return 0
	}
	var res int

	for i < len(words) && words[i][0] == s[0] {
		var j int
		for j < len(words[i]) && j < len(s) && words[i][j] == s[j] {
			j++
		}
		if j == len(words[i]) {
			res++
		}
		i++
	}
	return res
}
