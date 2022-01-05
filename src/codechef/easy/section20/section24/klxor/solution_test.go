package main

import "testing"

func runSample(t *testing.T, n, k int, S string, expect int) {
	res := solve(n, k, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 4
	S := "0110"
	expect := 2
	runSample(t, n, k, S, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	S := "1001"
	expect := 2
	runSample(t, n, k, S, expect)
}
