package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			ans++
			i++
			j++
			continue
		}

		j++
	}

	return ans
}
