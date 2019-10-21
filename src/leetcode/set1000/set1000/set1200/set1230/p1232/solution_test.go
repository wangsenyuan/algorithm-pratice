package p1232

import "testing"

func runSample(t *testing.T, coords [][]int, expect bool) {
	res := checkStraightLine(coords)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", coords, expect, res)
	}
}

func TestSample1(t *testing.T) {
	coords := [][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}
	expect := true
	runSample(t, coords, expect)
}
