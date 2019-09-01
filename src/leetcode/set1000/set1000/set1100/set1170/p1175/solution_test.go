package p1175

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := numPrimeArrangements(n)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 682289015)
}
