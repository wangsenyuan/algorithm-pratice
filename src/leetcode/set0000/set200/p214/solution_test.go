package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := shortestPalindrome(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aacecaaa", "aaacecaaa")
}

func TestSample2(t *testing.T) {
	runSample(t, "abcde", "edcbabcde")
}
