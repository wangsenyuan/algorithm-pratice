package p1895

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minOperationsToFlip(s)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1&(0|1)"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "(0&0)&(0&0&0)"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "(0|(1|0&1))"
	expect := 1
	runSample(t, s, expect)
}
