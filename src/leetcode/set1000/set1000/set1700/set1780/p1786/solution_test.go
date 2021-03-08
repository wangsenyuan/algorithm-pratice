package p1786

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := countRestrictedPaths(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10},
	}
	expect := 3
	runSample(t, n, E, expect)
}
