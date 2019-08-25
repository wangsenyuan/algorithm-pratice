package p1105

import "testing"

func runSample(t *testing.T, books [][]int, width int, expect int) {
	res := minHeightShelves(books, width)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", books, width, expect, res)
	}
}

func TestSample1(t *testing.T) {
	books := [][]int{
		{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2},
	}
	width := 4
	expect := 6
	runSample(t, books, width, expect)
}
