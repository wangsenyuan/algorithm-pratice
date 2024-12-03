package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 6, 2, 5}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 4, 4, 5, 1}
	expect := 11
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 8, 2, 3}
	expect := 10
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}
