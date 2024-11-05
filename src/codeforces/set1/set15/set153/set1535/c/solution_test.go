package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, bug got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0?10"
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "???"
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "?10??1100"
	expect := 25
	runSample(t, s, expect)
}
