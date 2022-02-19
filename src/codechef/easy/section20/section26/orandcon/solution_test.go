package main

import "testing"

func runSample(t *testing.T, x int) {
	res := solve(x)
	a, b, c := res[0], res[1], res[2]

	y := (a | b) & (b | c) & (c | a)

	if x != y || a == b || b == c || c == a {
		t.Errorf("Sample result %v not correct, expect %d, but got %d", res, x, y)
	}
}

func TestSaample1(t *testing.T) {
	for x := 1; x <= 100000; x++ {
		runSample(t, x)
	}
}
