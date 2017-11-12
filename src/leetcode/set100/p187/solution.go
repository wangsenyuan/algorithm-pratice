package main

import (
	"fmt"
	"sort"
)

func main() {
	repeat := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	for _, x := range repeat {
		fmt.Println(x)
	}
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	strs := make([]string, len(s)-9)

	for i := 0; i < len(s); i++ {
		j := i + 10
		if j > len(s) {
			break
		}
		strs[i] = s[i:j]
	}

	sort.Strings(strs)

	var res []string
	i := 0
	for i < len(strs) {
		j := i + 1
		for j < len(strs) && strs[j] == strs[i] {
			j++
		}
		if j-i > 1 {
			res = append(res, strs[i])
		}
		i = j
	}

	return res
}
