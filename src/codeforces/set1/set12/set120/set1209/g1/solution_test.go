package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a, nil)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 7, 3, 7, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1, 2, 3, 1, 1, 1, 50, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 6, 3, 3, 4, 4}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 3, 1, 3, 2, 1, 2}
	expect := 4
	runSample(t, a, expect)
}
