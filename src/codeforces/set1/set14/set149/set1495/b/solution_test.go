package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1, 2, 5, 4, 3}
	expect := 1
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 2, 4, 6, 5, 3, 7}
	expect := 0
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{1, 2, 4, 3, 5}
	expect := 0
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{3, 4, 5, 6, 7, 8, 9, 2, 1}
	expect := 0
	runSample(t, p, expect)
}
