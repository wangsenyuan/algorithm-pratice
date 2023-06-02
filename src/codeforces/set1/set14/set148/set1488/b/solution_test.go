package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "(()())((()))"
	k := 2
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "()()()"
	k := 3
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "(((())))"
	k := 1
	expect := 2
	runSample(t, s, k, expect)
}
