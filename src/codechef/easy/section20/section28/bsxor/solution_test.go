package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect bool) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	s := "111"
	expect := false
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 3
	s := "01100"
	expect := true
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 8, 3
	s := "01001111"
	expect := true
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 6, 2
	s := "000100"
	expect := false
	runSample(t, n, k, s, expect)
}
