package p1733

import "testing"

func runSample(t *testing.T, n int, langs [][]int, friends [][]int, expect int) {
	res := minimumTeachings(n, langs, friends)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	langs := [][]int{
		{2}, {1}, {2, 1}, {1}, {1, 2}, {1}, {2}, {1}, {1}, {2}, {1, 2}, {1, 2}, {1, 2}, {2, 1}, {1}, {2}, {1, 2},
	}
	frieds := [][]int{
		{15, 16}, {4, 13}, {3, 16}, {5, 14}, {1, 7}, {2, 11}, {3, 15}, {4, 16}, {7, 9}, {6, 13}, {6, 16}, {4, 10}, {6, 9}, {5, 6}, {7, 12}, {6, 12}, {3, 7}, {4, 7}, {8, 10},
	}
	expect := 3
	runSample(t, n, langs, frieds, expect)
}
