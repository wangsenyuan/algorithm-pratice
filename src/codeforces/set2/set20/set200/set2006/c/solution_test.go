package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21}
	expect := 18
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{70, 130, 90, 90, 90, 108, 612, 500, 451, 171, 193, 193}
	expect := 53
	runSample(t, a, expect)
}
