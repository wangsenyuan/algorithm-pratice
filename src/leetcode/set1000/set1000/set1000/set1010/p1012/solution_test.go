package p1012

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := bitwiseComplement(N)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 5)
}
