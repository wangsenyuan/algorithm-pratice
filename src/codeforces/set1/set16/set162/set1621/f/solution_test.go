package main

import "testing"

func runSample(t *testing.T, n, a, b, c int, s string, expect int) {
	res := int(solve(n, a, b, c, s))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b, c := 5, 2, 2, 1
	s := "01101"
	expect := 3
	runSample(t, n, a, b, c, s, expect)
}

func TestSample2(t *testing.T) {
	n, a, b, c := 6, 4, 3, 5
	s := "110001"
	expect := 11
	runSample(t, n, a, b, c, s, expect)
}

func TestSample3(t *testing.T) {
	n, a, b, c := 6, 3, 2, 1
	s := "011110"
	expect := 4
	runSample(t, n, a, b, c, s, expect)
}
