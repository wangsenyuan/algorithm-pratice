package p2810

import "testing"

func runSample(t *testing.T, items [][]int, k int, expect int64) {
	res := findMaximumElegance(items, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {10, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9}, {3, 10}, {3, 11},
	}
	k := 10
	var expect int64 = 137
	runSample(t, items, k, expect)
}
