package p878

import "testing"

func runSample(t *testing.T, N int, A int, B int, expect int) {
	res := nthMagicalNumber(N, A, B)
	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", N, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, A, B := 3, 6, 4
	expect := 8
	runSample(t, N, A, B, expect)
}
