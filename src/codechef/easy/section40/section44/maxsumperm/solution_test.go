package main

import "testing"

func runSample(t *testing.T, A []int, Q [][]int, expect int64) {
	sum, res := solve(A, Q)

	if sum != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
	n := len(A)
	pref := make([]int64, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + int64(res[i])
	}

	var x int64
	for _, cur := range Q {
		l, r := cur[0], cur[1]
		x += pref[r] - pref[l-1]
	}

	if x != expect {
		t.Fatalf("Sample result %v, getting %d, not %d", res, x, expect)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{1, 4},
		{2, 3},
	}
	var expect int64 = 23
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1}
	Q := [][]int{
		{1, 1},
		{1, 2},
		{2, 2},
	}
	var expect int64 = 4
	runSample(t, A, Q, expect)
}
