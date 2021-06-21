package main

import "testing"

func runSample(t *testing.T, n int, S string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	S := "110011"
	expect := 2
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := "0001"
	expect := 0
	runSample(t, n, S, expect)
}
