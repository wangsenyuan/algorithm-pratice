package main

import (
	"testing"
)

func runSample(t *testing.T, n int, S string, expect int64) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := "RD"
	var expect int64 = 13
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	S := "DRDRDRDR"
	var expect int64 = 9
	runSample(t, n, S, expect)
}
