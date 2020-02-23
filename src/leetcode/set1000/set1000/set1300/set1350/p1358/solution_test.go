package p1358

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := numberOfSubstrings(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcabc", 10)
}

func TestSample2(t *testing.T) {
	runSample(t, "aaacb", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "abc", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, "acbbcac", 11)
}
