package p3380

import "testing"

func runSample(t *testing.T, xs []int, ys []int, expect int) {
	res := maxRectangleArea(xs, ys)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	xCoord := []int{1, 1, 3, 3}
	yCoord := []int{1, 3, 1, 3}
	expect := 4
	runSample(t, xCoord, yCoord, expect)
}

func TestSample2(t *testing.T) {
	xCoord := []int{1, 1, 3, 3, 2}
	yCoord := []int{1, 3, 1, 3, 2}
	expect := -1
	runSample(t, xCoord, yCoord, expect)
}

func TestSample3(t *testing.T) {
	xCoord := []int{1, 1, 3, 3, 1, 3}
	yCoord := []int{1, 3, 1, 3, 2, 2}
	expect := 2
	runSample(t, xCoord, yCoord, expect)
}

func TestSample4(t *testing.T) {
	xCoord := []int{89, 55, 89, 55, 0, 34, 17, 71, 98, 90, 63, 49, 76, 72, 4, 46, 67, 94, 52, 6}
	yCoord := []int{58, 69, 69, 58, 100, 36, 14, 40, 13, 41, 29, 23, 47, 52, 95, 49, 37, 77, 54, 59}
	expect := 374
	runSample(t, xCoord, yCoord, expect)
}
