package main

import "testing"

func runSample(t *testing.T, n, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 6, 3
	s := "abcbac"
	expect := 6
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 6
	s := "aaa"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 7, 1000
	s := "abczgyo"
	expect := 5
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 5, 4
	s := "ababa"
	expect := 4
	runSample(t, n, k, s, expect)
}

func TestSample5(t *testing.T) {
	n, k := 20, 10
	s := "aaebdbabdbbddaadaadc"
	expect := 15
	runSample(t, n, k, s, expect)
}

func TestSample6(t *testing.T) {
	n, k := 20, 5
	s := "ecbedececacbcbccbdec"
	expect := 10
	runSample(t, n, k, s, expect)
}
