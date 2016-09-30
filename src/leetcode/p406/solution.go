package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}
	fmt.Println(reconstructQueue(people))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(byHeight(people))

	var res [][]int

	for _, person := range people {
		_, k := person[0], person[1]
		res = insertAt(res, k, person)
	}

	return res
}

func insertAt(xs [][]int, k int, x []int) [][]int {
	a := make([][]int, k+1)
	copy(a, xs[:k])
	a[k] = x
	return append(a, xs[k:]...)
}

type byHeight [][]int

func (p byHeight) Len() int {
	return len(p)
}

func (p byHeight) Less(i, j int) bool {
	a, b := p[i], p[j]
	if a[0] != b[0] {
		return b[0]-a[0] < 0
	}
	return a[1]-b[1] < 0
}

func (p byHeight) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
