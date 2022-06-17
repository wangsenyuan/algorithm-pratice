package main

import "testing"

func runSample(t *testing.T, A, B int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B := 12, 8
	expect := 3
	runSample(t, A, B, expect)
}
