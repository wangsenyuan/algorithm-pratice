package p851

import "testing"

func runSample(t *testing.T, rects [][]int, expect int) {
	res := rectangleArea(rects)
	if res != expect {
		t.Errorf("Sample %v, expect %v, but got %v", rects, expect, res)
	}
}

func TestSample1(t *testing.T) {
	rects := [][]int{
		{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1},
	}
	expect := 6

	runSample(t, rects, expect)
}

func TestSample2(t *testing.T) {
	rects := [][]int{
		{0, 0, 1000000000, 1000000000},
	}
	expect := 49

	runSample(t, rects, expect)
}

func TestSample3(t *testing.T) {
	rects := [][]int{
		{49, 40, 62, 100}, {11, 83, 31, 99}, {19, 39, 30, 99},
	}
	expect := 1584

	runSample(t, rects, expect)
}
