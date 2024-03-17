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
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 5, 5}
	expect := 6
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	expect := 540
	runSample(t, a, expect)
}
