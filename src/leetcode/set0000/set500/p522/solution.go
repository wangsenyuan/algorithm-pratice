package main

import (
	"fmt"
)

func main() {

	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "cb"}))

}

func findLUSlength(strs []string) int {
	if len(strs) == 0 {
		return -1
	}
	if len(strs) == 1 {
		return len(strs[0])
	}

	var res = -1
	for i, cur := range strs {
		j := 0
		for j < len(strs) {
			if i == j {
				j++
				continue
			}
			if isSubSeq(cur, strs[j]) {
				break
			}
			j++
		}
		if j == len(strs) && len(cur) > res {
			res = len(cur)
		}
	}
	return res
}

func isSubSeq(x, y string) bool {
	i, j := 0, 0
	for i < len(x) && j < len(y) {
		if x[i] == y[j] {
			i++
		}
		j++
	}
	return i == len(x)
}
