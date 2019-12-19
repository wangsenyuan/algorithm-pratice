package main

import "testing"

func runSample(t *testing.T, n int, A, B string, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %s %s, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := "00"
	B := "01"
	expect := 2
	runSample(t, n, A, B, expect)
}
