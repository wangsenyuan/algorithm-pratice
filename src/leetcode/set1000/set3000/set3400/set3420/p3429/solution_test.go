package p3429

import "testing"

func runSample(t *testing.T, n int, cost [][]int, expect int) {
	res := minCost(n, cost)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	cost := [][]int{
		{3, 5, 7}, {6, 2, 9}, {4, 8, 1}, {7, 3, 5},
	}
	expect := 9
	runSample(t, n, cost, expect)
}
