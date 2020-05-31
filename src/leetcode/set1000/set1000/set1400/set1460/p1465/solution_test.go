package p1465

import "testing"

func runSample(t *testing.T, h, w int, horizontal, vertical []int, expect int) {
	res := maxArea(h, w, horizontal, vertical)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", h, w, horizontal, vertical, expect, res)
	}
}

func TestSample1(t *testing.T) {
	h, w := 5, 4
	horizontal := []int{1, 2, 4}
	vertical := []int{1, 3}
	runSample(t, h, w, horizontal, vertical, 4)
}
