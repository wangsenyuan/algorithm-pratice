package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 19
	s := "Hello! Do you like fish? Why?"
	expect := 3
	runSample(t, n, s, expect)
}
