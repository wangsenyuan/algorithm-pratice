package b

import "testing"

func runSample(t *testing.T, n int, relation [][]int, k int, expect int) {
	res := numWays(n, relation, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	k := 3
	runSample(t, n, relation, k, 3)
}

func TestSample2(t *testing.T) {
	n := 3
	relation := [][]int{{0, 2}, {2, 1}}
	k := 2
	runSample(t, n, relation, k, 0)
}
