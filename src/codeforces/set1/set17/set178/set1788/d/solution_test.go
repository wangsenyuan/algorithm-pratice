package main

import "testing"

func runSample(t *testing.T, x []int, expect int) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{1, 2, 4, 6}
	expect := 11
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1, 3, 5, 11, 15}
	expect := 30
	runSample(t, x, expect)
}
