package main

import "testing"

func runSample(t *testing.T, n, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 2
	s := "1100110"
	expect := 16
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 1
	s := "10001000"
	expect := 10
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 0
	s := "10010"
	expect := 1
	runSample(t, n, k, s, expect)
}
