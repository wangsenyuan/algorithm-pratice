package p2121

import "testing"

func runSample(t *testing.T, s string, locked string, expect bool) {
	res := canBeValid(s, locked)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "))()))"
	locked := "010100"
	expect := true
	runSample(t, s, locked, expect)
}

func TestSample2(t *testing.T) {
	s := ")"
	locked := "0"
	expect := false
	runSample(t, s, locked, expect)
}
