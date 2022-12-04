package p2493

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := magnificentSets(n, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6},
	}
	expect := 4
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2}, {1, 3}, {3, 2},
	}
	expect := -1
	runSample(t, n, edges, expect)
}
