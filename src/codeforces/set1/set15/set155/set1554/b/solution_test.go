package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	a := []int{1, 1, 3}
	expect := -1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{1, 2}
	expect := -4
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	a := []int{0, 1, 2, 3}
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	k := 6
	a := []int{3, 2, 0, 0, 5, 6}
	expect := 12
	runSample(t, a, k, expect)
}
