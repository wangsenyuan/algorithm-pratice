package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 1, 3}
	expect := 9
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 5, 4, 3, 2}
	expect := 17
	runSample(t, p, expect)
}
