package p793

import "testing"

func runSample(t *testing.T, K int, expect int) {
	res := preimageSizeFZF(K)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 0)
}
