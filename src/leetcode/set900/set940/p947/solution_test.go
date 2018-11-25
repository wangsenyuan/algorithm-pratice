package p947

import "testing"

func runSample(t *testing.T, stones [][]int, expect int) {
	res := removeStones(stones)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := [][]int{
		{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2},
	}
	expect := 5
	runSample(t, stones, expect)
}

func TestSample2(t *testing.T) {
	stones := [][]int{
		{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2},
	}
	expect := 3
	runSample(t, stones, expect)
}
func TestSample3(t *testing.T) {
	stones := [][]int{
		{0, 0},
	}
	expect := 0
	runSample(t, stones, expect)
}
