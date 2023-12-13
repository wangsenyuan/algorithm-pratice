package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abacaba"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "codeforces"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "oooooooo"
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abcdef"
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "mewheniseearulhiiarul"
	expect := 4
	runSample(t, s, expect)
}
