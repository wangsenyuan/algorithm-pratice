package main

import "testing"

func runSample(t *testing.T, A, B int, expect int) {
	res := solve(A, B)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B := 2, 4
	expect := 24
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A, B := 8, 3
	expect := 178
	runSample(t, A, B, expect)
}
