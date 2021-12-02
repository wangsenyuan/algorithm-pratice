package main

import "testing"

func runSample(t *testing.T, n, m int, S string, expect int) {
	res := solve(n, m, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 2
	S := "ABACA"
	expect := 3
	runSample(t, n, m, S, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 1
	S := "ABAB"
	expect := 1
	runSample(t, n, m, S, expect)
}

func TestSample3(t *testing.T) {
	n, m := 13, 4
	S := "AABAACABDAAAE"
	expect := 8
	runSample(t, n, m, S, expect)
}
