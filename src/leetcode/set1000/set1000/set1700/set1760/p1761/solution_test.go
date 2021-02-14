package p1761

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := minTrioDegree(n, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6},
	}
	expect := 0
	runSample(t, n, edges, expect)
}
