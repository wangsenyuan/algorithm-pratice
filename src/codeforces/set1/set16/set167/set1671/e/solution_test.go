package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	s := "BAAAAAAAABBABAB"
	expect := 16
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	s := "BAA"
	expect := 1
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	s := "AAB"
	expect := 2
	runSample(t, n, s, expect)
}
