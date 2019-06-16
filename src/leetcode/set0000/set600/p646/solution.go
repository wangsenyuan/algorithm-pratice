package main

import (
	"sort"
	"fmt"
)

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}

	sort.Sort(ByEnd(pairs))

	res := 1

	j := 0

	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] <= pairs[j][1] {
			continue
		}
		res ++
		j = i
	}

	return res
}

type ByEnd [][]int

func (this ByEnd) Len() int {
	return len(this)
}

func (this ByEnd) Less(i, j int) bool {
	a, b := this[i], this[j]
	return a[1] < b[1] || (a[1] == b[1] && a[0] > b[0])
}

func (this ByEnd) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println(findLongestChain(pairs))
}
