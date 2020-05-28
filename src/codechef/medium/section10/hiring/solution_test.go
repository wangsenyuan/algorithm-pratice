package main

import (
	"testing"
)

func runSample(t *testing.T, n int, A []string, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []string{
		"10",
		"01",
		"11",
		"10",
	}
	expect := 7
	runSample(t, n, A, expect)
}
