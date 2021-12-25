package main

import "testing"

func runSample(t *testing.T, n, c int, s string, expect bool) {
	res := solve(n, c, s)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, c := 4, 1
	s := "1100"
	expect := true
	runSample(t, n, c, s, expect)
}

func TestSample2(t *testing.T) {
	n, c := 4, 0
	s := "0101"
	expect := false
	runSample(t, n, c, s, expect)
}
