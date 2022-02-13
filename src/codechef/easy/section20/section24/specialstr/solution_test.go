package main

import "testing"

func runSample(t *testing.T, n int, s string, m int, expect int) {
	res := solve(n, s, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	s := "ababbdced"
	m := 4
	expect := 3
	runSample(t, n, s, m, expect)
}
