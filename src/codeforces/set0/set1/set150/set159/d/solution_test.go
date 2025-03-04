package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aa"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaa"
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abacaba"
	expect := 36
	runSample(t, s, expect)
}
