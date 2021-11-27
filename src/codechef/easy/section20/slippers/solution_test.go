package main

import "testing"

func runSample(t *testing.T, S []int, expect int) {
	res := solve(S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []int{38, 43, 41}
	expect := 6
	runSample(t, S, expect)
}
