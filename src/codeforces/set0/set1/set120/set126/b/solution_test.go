package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := "fixprefixsuffix"
	expect := "fix"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcdabc"
	expect := "Just a legend"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "qwertyqwertyqwerty"
	expect := "qwerty"
	runSample(t, s, expect)
}
