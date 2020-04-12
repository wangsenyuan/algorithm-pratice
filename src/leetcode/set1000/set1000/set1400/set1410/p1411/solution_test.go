package p1411

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := numOfWays(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 54)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 106494)
}

func TestSample4(t *testing.T) {
	runSample(t, 5000, 30228214)
}
