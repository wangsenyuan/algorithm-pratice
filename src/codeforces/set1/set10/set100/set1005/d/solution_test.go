package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3121`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1000000000000000000000000000000000`
	expect := 33
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `201920181`
	expect := 4
	runSample(t, s, expect)
}
