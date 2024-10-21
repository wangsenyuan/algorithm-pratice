package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{50}
	expect := 2
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{10, 20, 50}
	expect := 112
	runSample(t, p, expect)
}
