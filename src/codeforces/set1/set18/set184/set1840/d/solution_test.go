package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1, 7, 7, 9, 9, 9}
	expect := 0
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{5, 4, 2, 1, 30, 60}
	expect := 2
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{14, 19, 37, 59, 1, 4, 4, 98, 73}
	expect := 13
	runSample(t, p, expect)
}
