package main

import "testing"

func runSample(t *testing.T, n, m int, trips [][]int, expect bool) {
	res := solve(n, m, trips)

	if expect != (len(res) > 0) {
		t.Errorf("Sample expect %t, but got %v", expect, res)
		return
	}

	if !expect {
		return
	}

	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + res[i]
	}

	for _, cur := range trips {
		x, y, z := cur[0], cur[1], cur[2]
		if sum[y]-sum[x-1] != int64(z) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	trips := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 4, -2},
		{1, 4, 1},
	}
	expect := true

	runSample(t, n, m, trips, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 3
	trips := [][]int{
		{1, 3, 4},
		{4, 5, 6},
		{1, 5, 9},
	}
	expect := false

	runSample(t, n, m, trips, expect)
}
