package p935

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := knightDialer(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 20)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 46)
}
