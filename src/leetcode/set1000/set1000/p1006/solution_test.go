package p1006

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := clumsy(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 12)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 6)
}
