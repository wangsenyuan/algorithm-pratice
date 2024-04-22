package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 1, 22, 14}
	expect := 13
	runSample(t, a, expect)
}
