package main

import "testing"

func runSample(t *testing.T, A string, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "11011"
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := "0111010"
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := "1100"
	expect := 1
	runSample(t, A, expect)
}
