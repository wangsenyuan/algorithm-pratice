package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "?01???"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "**12"
	expect := 0
	runSample(t, s, expect)
}
