package p044

import "testing"

func runSample(t *testing.T, s, p string, expect bool) {
	res := isMatch(s, p)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", s, p, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aa", "a", false)
}

func TestSample2(t *testing.T) {
	runSample(t, "aa", "*", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "cb", "?a", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "adceb", "*a*b", true)
}

func TestSample5(t *testing.T) {
	runSample(t, "acdcb", "a*c?b", false)
}

func TestSample6(t *testing.T) {
	runSample(t, "ho", "**ho", true)
}
