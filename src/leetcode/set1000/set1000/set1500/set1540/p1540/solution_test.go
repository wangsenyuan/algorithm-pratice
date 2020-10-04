package p1540

import "testing"

func runSample(tc *testing.T, s, t string, k int, expect bool) {
	res := canConvertString(s, t, k)

	if res != expect {
		tc.Errorf("Sample %s %s %d, expect %t, but got %t", s, t, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "input", "ouput", 9, true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abc", "bcd", 10, false)
}

func TestSample3(t *testing.T) {
	runSample(t, "aab", "bbb", 27, true)
}

func TestSample4(t *testing.T) {
	runSample(t, "iqssxdlb", "dyuqrwyr", 40, true)
}
