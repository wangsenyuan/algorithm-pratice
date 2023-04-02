package p2608

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := findShortestCycle(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1}, {0, 2},
	}
	expect := -1
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 13
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 0}, {0, 3}, {3, 4}, {4, 5}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}, {11, 12}, {12, 0}, {2, 7}, {2, 4}, {1, 8}, {1, 11},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	edges := [][]int{
		{4, 1}, {5, 1}, {3, 2}, {5, 0}, {4, 0}, {3, 0}, {2, 1},
	}
	expect := 4
	runSample(t, n, edges, expect)
}
