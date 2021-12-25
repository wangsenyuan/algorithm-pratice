package main

import "testing"

func runSample(t *testing.T, S string, expect int64) {
	res := solve(len(S), S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "276", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "762", 6)
}
