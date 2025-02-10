package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 2
	expect := 12
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 11, 3
	expect := 18
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 55, 13
	expect := 196
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 5801, 6
	expect := 1975581
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 8919, 64
	expect := 958900
	runSample(t, n, k, expect)
}

func TestSample6(t *testing.T) {
	n, k := 8765432, 1
	expect := 38416403456028
	runSample(t, n, k, expect)
}
