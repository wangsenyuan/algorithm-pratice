package p1027

import "testing"

func runSample(t *testing.T, clips [][]int, X int, expect int) {
	res := videoStitching(clips, X)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", clips, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	X := 10
	expect := 3
	runSample(t, clips, X, expect)
}

func TestSample2(t *testing.T) {
	clips := [][]int{{0, 1}, {1, 2}}
	X := 5
	expect := -1
	runSample(t, clips, X, expect)
}

func TestSample3(t *testing.T) {
	clips := [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	X := 9
	expect := 3
	runSample(t, clips, X, expect)
}
