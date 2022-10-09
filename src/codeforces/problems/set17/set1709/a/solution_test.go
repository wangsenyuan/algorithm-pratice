package main

import "testing"

func runSample(t *testing.T, x int, A []int, expect bool) {
	res := solve(x, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 3
	A := []int{0, 1, 2}
	expect := true

	runSample(t, x, A, expect)
}

func TestSample2(t *testing.T) {
	x := 1
	A := []int{0, 3, 2}
	expect := false

	runSample(t, x, A, expect)
}
