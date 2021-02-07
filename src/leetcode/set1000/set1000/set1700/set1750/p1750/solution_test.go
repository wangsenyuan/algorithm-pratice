package p1750

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minimumLength(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ca"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "cabaabac"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aabccabba"
	expect := 3
	runSample(t, s, expect)
}
