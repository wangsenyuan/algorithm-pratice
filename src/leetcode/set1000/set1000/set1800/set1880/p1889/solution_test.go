package p1889

import "testing"

func runSample(t *testing.T, packages []int, boxes [][]int, expect int) {
	res := minWastedSpace(packages, boxes)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	packages := []int{2, 3, 5}
	boxes := [][]int{
		{4, 8}, {2, 8},
	}

	expect := 6
	runSample(t, packages, boxes, expect)
}
