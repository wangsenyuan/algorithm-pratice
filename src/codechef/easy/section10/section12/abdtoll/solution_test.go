package main

import "testing"

func runSample(t *testing.T, n, x int, k int64, tl []int, edges [][]int, expect int64) {
	res := solve(n, x, k, tl, edges)

	if res != expect {
		t.Errorf("sample %d %d %d %v %v, expect %d, but got %d", n, x, k, tl, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 4, 2
	k := int64(6)

	tl := []int{4, 7, 6, 2}
	edges := [][]int{
		{2, 1}, {2, 3}, {2, 4},
	}

	expect := int64(14)

	runSample(t, n, x, k, tl, edges, expect)
}

func TestSample2(t *testing.T) {
	n, x := 3, 1
	k := int64(8)

	tl := []int{2, 3, 1}
	edges := [][]int{
		{2, 1}, {2, 3},
	}

	expect := int64(0)

	runSample(t, n, x, k, tl, edges, expect)
}
