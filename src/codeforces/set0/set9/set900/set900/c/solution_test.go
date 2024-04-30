package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1}
	expect := 1
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{5, 1, 2, 3, 4}
	expect := 5
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{2, 1, 5, 3, 4}
	expect := 5
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{2, 1, 4, 5, 3}
	expect := 1
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{3, 1, 2, 5, 4}
	expect := 3
	runSample(t, p, expect)
}
