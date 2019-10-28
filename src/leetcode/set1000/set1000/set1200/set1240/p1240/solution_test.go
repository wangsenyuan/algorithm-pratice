package p1240

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := tilingRectangle(n, m)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 8, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 13, 6)
}
