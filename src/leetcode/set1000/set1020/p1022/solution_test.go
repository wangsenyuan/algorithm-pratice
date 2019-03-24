package p1022

import "testing"

func runSample(t *testing.T, K int, expect int) {
	res := smallestRepunitDivByK(K)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3)
}
