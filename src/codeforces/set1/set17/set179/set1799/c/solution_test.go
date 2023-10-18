package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a"
	expect := "a"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aab"
	expect := "aba"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abb"
	expect := "bab"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abc"
	expect := "bca"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "aabb"
	expect := "abba"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "aabbb"
	expect := "abbba"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "aaabb"
	expect := "ababa"
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "abbb"
	expect := "bbab"
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "abbbb"
	expect := "bbabb"
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := "abbcc"
	expect := "bbcca"
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := "eaga"
	expect := "agea"
	runSample(t, s, expect)
}

func TestSample12(t *testing.T) {
	s := "ffcaba"
	expect := "acffba"
	runSample(t, s, expect)
}
