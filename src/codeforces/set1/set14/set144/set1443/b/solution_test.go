package main

import "testing"

func runSample(t *testing.T, s string, a, b int, expect int) {
	res := solve(s, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01000010"
	a, b := 1, 1
	expect := 2
	runSample(t, s, a, b, expect)
}

func TestSample2(t *testing.T) {
	s := "01101110"
	a, b := 5, 1
	expect := 6
	runSample(t, s, a, b, expect)
}
