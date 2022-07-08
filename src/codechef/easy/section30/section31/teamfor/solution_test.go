package main

import "testing"

func runSample(t *testing.T, n int, S string, T string, expect int) {
	res := solve(n, S, T)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := "101"
	T := "001"
	expect := 1
	runSample(t, n, S, T, expect)
}
