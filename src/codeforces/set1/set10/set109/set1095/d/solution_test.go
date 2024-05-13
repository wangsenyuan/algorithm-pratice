package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a [][]int) {
	res := solve(len(a), a)
	n := len(a)

	pos := make([]int, n+1)

	for i := 0; i < n; i++ {
		pos[res[i]] = i
	}

	for i := 1; i <= n; i++ {
		cur := make([]int, 2)
		cur[0] = res[(pos[i]+1)%n]
		cur[1] = res[(pos[i]+2)%n]
		tmp := make([]int, 2)
		copy(tmp, a[i-1])
		sort.Ints(tmp)
		sort.Ints(cur)

		if tmp[0] != cur[0] || tmp[1] != cur[1] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{3, 5},
		{1, 4},
		{2, 4},
		{1, 5},
		{2, 3},
	}
	runSample(t, a)
}
