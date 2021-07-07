package p1927

import "testing"

func runSample(t *testing.T, points [][]int, k int, expect int) {
	res := minDayskVariants(points, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{1, 1}, {6, 1}}
	k := 2
	expect := 3
	runSample(t, points, k, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{3, 3}, {1, 2}, {9, 2}}
	k := 2
	expect := 2
	runSample(t, points, k, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{{3, 3}, {1, 2}, {9, 2}}
	k := 3
	expect := 4
	runSample(t, points, k, expect)
}
