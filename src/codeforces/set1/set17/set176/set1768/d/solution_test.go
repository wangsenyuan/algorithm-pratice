package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 1}
	expect := 0
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 2}
	expect := 1
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{3, 4, 1, 2}
	expect := 3
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{2, 4, 3, 1}
	expect := 1
	runSample(t, p, expect)
}
