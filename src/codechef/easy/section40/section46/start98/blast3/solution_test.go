package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaaa"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcde"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aabcd"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abcad"
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "abcadcae"
	expect := false
	runSample(t, s, expect)
}
