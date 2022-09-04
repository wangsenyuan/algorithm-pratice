package p2397

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := isStrictlyPalindromic(n)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, false)
}
