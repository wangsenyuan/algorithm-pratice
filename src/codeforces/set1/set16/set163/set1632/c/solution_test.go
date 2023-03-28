package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 1, 3
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 5, 8
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 2, 5
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a, b := 3, 19
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a, b := 56678, 164422
	expect := 23329
	runSample(t, a, b, expect)
}
