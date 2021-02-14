package p1759

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countHomogenous(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abbcccaa"
	expect := 13
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "zzzzz"
	expect := 15
	runSample(t, s, expect)
}
