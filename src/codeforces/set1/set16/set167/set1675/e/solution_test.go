package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cba"
	k := 2
	expect := "aaa"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "fgde"
	k := 5
	expect := "agaa"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "gndcafb"
	k := 5
	expect := "bnbbabb"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "ekyv"
	k := 19
	expect := "aapp"
	runSample(t, s, k, expect)
}
