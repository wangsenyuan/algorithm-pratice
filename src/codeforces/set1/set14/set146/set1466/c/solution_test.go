package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSmaple1(t *testing.T) {
	s := "babba"
	expect := 1
	runSample(t, s, expect)
}

func TestSmaple2(t *testing.T) {
	s := "abaac"
	expect := 1
	runSample(t, s, expect)
}

func TestSmaple3(t *testing.T) {
	s := "codeforces"
	expect := 0
	runSample(t, s, expect)
}

func TestSmaple4(t *testing.T) {
	s := "zeroorez"
	expect := 1
	runSample(t, s, expect)
}

func TestSmaple5(t *testing.T) {
	s := "abcdcba"
	expect := 1
	runSample(t, s, expect)
}

func TestSmaple6(t *testing.T) {
	s := "bbbbbbb"
	expect := 4
	runSample(t, s, expect)
}
