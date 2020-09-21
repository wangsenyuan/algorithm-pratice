package p1594

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := maxUniqueSplit(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababccc"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aa"
	expect := 1
	runSample(t, s, expect)
}
