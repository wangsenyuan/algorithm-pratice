package p2492

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect int) {
	res := minScore(n, roads)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7},
	}
	expect := 5
	runSample(t, n, roads, expect)
}
