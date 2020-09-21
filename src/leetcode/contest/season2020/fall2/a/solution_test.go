package a

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := paintingPlan(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 24, 630)
}
