package p1592

import "testing"

func runSample(t *testing.T, targetRect [][]int, expect bool) {
	res := isPrintable(targetRect)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1}, {1, 2, 2, 1}, {1, 2, 2, 1}, {1, 1, 1, 1}}
	expect := true
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1}, {1, 1, 3, 3}, {1, 1, 3, 4}, {5, 5, 1, 4}}
	expect := true
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{{1, 2, 1}, {2, 1, 2}, {1, 2, 1}}
	expect := false
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{{1, 1, 1}, {3, 1, 3}}
	expect := false
	runSample(t, grid, expect)
}
