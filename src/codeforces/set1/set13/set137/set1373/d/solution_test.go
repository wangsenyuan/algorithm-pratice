package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 7, 3, 4, 7, 6, 2, 9}
	expect := 26
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1, 2, 1}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 8, 4, 5, 7, 6, 8, 9, 7, 3}
	expect := 37
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 1, 2, 1}
	expect := 5
	runSample(t, a, expect)
}
