package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{6, 5, 7, 8, 4, 3, 1, 2}
	expect := 4
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{3, 1, 4, 2}
	expect := -1
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{7, 8, 4, 3, 1, 2, 6, 5}
	expect := -1
	runSample(t, p, expect)
}
