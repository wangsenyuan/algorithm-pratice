package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int) {
	res := solve(n, pairs)

	left := make([]int, n)
	right := make([]int, n)
	for i, x := range res {
		left[x-1] = i
		right[x-1] = i
	}

	set := NewDSU(n)

	for _, cur := range pairs {
		x, y := cur[0]-1, cur[1]-1
		x = set.Find(x)
		y = set.Find(y)
		if left[x] > left[y] {
			x, y = y, x
		}
		if right[x]+1 != left[y] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		set.Union(x, y)
		p := set.Find(x)
		left[p] = left[x]
		right[p] = right[y]
	}
}

func TestSample1(t *testing.T) {
	n := 5
	pairs := [][]int{
		{1, 4},
		{2, 5},
		{3, 1},
		{4, 5},
	}
	runSample(t, n, pairs)
}
