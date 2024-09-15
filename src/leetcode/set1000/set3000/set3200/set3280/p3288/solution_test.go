package p3288

import "testing"

func runSample(t *testing.T, coords [][]int, k int, expect int) {
	res := maxPathLength(coords, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	coords := [][]int{
		{3, 1}, {2, 2}, {4, 1}, {0, 0}, {5, 3},
	}
	k := 1
	expect := 3
	runSample(t, coords, k, expect)
}

func TestSample2(t *testing.T) {
	coords := [][]int{
		{2, 1}, {7, 0}, {5, 6},
	}
	k := 2
	expect := 2
	runSample(t, coords, k, expect)
}
