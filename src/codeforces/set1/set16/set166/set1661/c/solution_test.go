package main

import "testing"

func runSample(t *testing.T, h []int, expect int) {
	res := solve(h)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 4}
	expect := 4
	runSample(t, a, expect)
}
