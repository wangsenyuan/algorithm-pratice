package p1830

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := makeStringSorted(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cba"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aabaa"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "cdbea"
	expect := 63
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "leetcodeleetcodeleetcode"
	expect := 982157772
	runSample(t, s, expect)
}
