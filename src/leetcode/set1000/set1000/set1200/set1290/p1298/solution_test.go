package p1298

import "testing"

func runSample(t *testing.T, status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int, expect int) {
	res := maxCandies(status, candies, keys, containedBoxes, initialBoxes)

	if res != expect {
		t.Errorf("Sample %v %v %v %v %v, expect %d, but got %d", status, candies, keys, containedBoxes, initialBoxes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	status := []int{1, 0, 1, 0}
	candies := []int{7, 5, 4, 100}
	keys := [][]int{
		{}, {}, {1}, {},
	}
	containedBoxes := [][]int{
		{1, 2},
		{3},
		{},
		{},
	}
	initialBoxes := []int{0}
	expect := 16
	runSample(t, status, candies, keys, containedBoxes, initialBoxes, expect)
}

func TestSample2(t *testing.T) {
	status := []int{1, 1, 1}
	candies := []int{2, 3, 2}
	keys := [][]int{
		{}, {}, {},
	}
	containedBoxes := [][]int{
		{},
		{},
		{},
	}
	initialBoxes := []int{2, 1, 0}
	expect := 7
	runSample(t, status, candies, keys, containedBoxes, initialBoxes, expect)
}
