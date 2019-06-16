package p741

import "testing"

func TestSample(t *testing.T) {
	grid := [][]int{
		[]int{0, 1, -1},
		[]int{1, 0, -1},
		[]int{1, 1, 1},
	}
	res := cherryPickup(grid)
	if res != 5 {
		t.Errorf("the sample %v should give answer 5, but got %d", grid, res)
	}
}
