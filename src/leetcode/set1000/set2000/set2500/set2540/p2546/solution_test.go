package p2546

import "testing"

func runSample(t *testing.T, S string, D string, expect bool) {
	res := makeStringsEqual(S, D)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1010"
	target := "0110"
	expect := true
	runSample(t, s, target, expect)
}

func TestSample2(t *testing.T) {
	s := "11"
	target := "00"
	expect := false
	runSample(t, s, target, expect)
}
