package main

import "testing"

func runSample(t *testing.T, n int, P []int, Q [][]int, expect bool) {
	res := solve(n, P, Q)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	ord := make([]int, n)
	for i := 0; i < n; i++ {
		ord[res[i]-1] = i
	}

	for i := 0; i < n; i++ {
		j := P[i]
		j--
		if j >= 0 {
			if ord[j] > ord[i] {
				t.Fatalf("Sample result %v not correct, as %d needs to preceeds %d, but not", res, j, i)
			}
		}
	}

	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = -1
	}
	for _, cur := range Q {
		x, y := cur[0], cur[1]
		next[x-1] = y - 1
	}

	for i := 0; i+1 < n; i++ {
		x := res[i] - 1
		if next[x] != -1 && next[x] != res[i+1]-1 {
			t.Fatalf("Sample result %v not correct, as condition (%d %d) not meet", res, x, next[x])
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	P := []int{2, 3, 0, 5, 3}
	Q := [][]int{
		{1, 5},
		{5, 4},
	}

	expect := true
	runSample(t, n, P, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	P := []int{2, 3, 0, 5, 3}
	Q := [][]int{
		{1, 5},
		{5, 1},
	}

	expect := false
	runSample(t, n, P, Q, expect)
}
