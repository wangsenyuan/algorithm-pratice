package main

import "testing"

func runSample(t *testing.T, S string, X, Y int, expect int64) {
	res := solve(len(S), X, Y, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X, Y := 2, 1
	S := "SENWS"
	var expect int64 = 4
	runSample(t, S, X, Y, expect)
}
