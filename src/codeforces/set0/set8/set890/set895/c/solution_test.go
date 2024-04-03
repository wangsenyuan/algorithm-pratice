package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := 15
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 2, 2}
	expect := 7
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 4, 5, 8}
	expect := 7
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	expect := 7
	runSample(t, a, expect)
}
