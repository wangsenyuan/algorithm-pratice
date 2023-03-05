package p2581

import "testing"

func runSample(t *testing.T, edges [][]int, guesses [][]int, k int, expect int) {
	res := rootCount(edges, guesses, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3}, {4, 2},
	}
	guesses := [][]int{
		{1, 3}, {0, 1}, {1, 0}, {2, 4},
	}
	k := 3
	expect := 3
	runSample(t, edges, guesses, k, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 4},
	}
	guesses := [][]int{
		{1, 0}, {3, 4}, {2, 1}, {3, 2},
	}
	k := 1
	expect := 5
	runSample(t, edges, guesses, k, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{0, 1}, {2, 0}, {0, 3}, {4, 2}, {3, 5}, {6, 0}, {1, 7}, {2, 8}, {2, 9}, {4, 10}, {9, 11}, {3, 12}, {13, 8}, {14, 9}, {15, 9}, {10, 16},
	}
	guesses := [][]int{
		{8, 2}, {12, 3}, {0, 1}, {16, 10},
	}
	k := 2
	expect := 4
	runSample(t, edges, guesses, k, expect)
}
