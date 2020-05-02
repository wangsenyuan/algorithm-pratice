package p1429

import "testing"

func runSample(t *testing.T, hats [][]int, expect int) {
	res := numberWays(hats)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", hats, expect, res)
	}
}

func TestSample1(t *testing.T) {
	hats := [][]int{
		{3, 4}, {4, 5}, {5},
	}
	runSample(t, hats, 1)
}

func TestSample2(t *testing.T) {
	hats := [][]int{
		{3, 5, 1}, {3, 5},
	}
	runSample(t, hats, 4)
}

func TestSample3(t *testing.T) {
	hats := [][]int{
		{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4},
	}
	runSample(t, hats, 24)
}

func TestSample4(t *testing.T) {
	hats := [][]int{
		{1, 2, 3}, {2, 3, 5, 6}, {1, 3, 7, 9}, {1, 8, 9}, {2, 5, 7},
	}
	runSample(t, hats, 111)
}
