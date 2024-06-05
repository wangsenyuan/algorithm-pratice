package main

import "testing"

func runSample(t *testing.T, a []int, k int, l int, expect int) {
	res := solve(a, k, l)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, l := 2, 1
	a := []int{2, 2, 1, 2, 3, 2, 2, 3}
	expect := 7
	runSample(t, a, k, l, expect)
}

func TestSample2(t *testing.T) {
	k, l := 1, 0
	a := []int{10, 10}
	expect := 20
	runSample(t, a, k, l, expect)
}

func TestSample3(t *testing.T) {
	k, l := 2, 1
	a := []int{5, 2}
	expect := 2
	runSample(t, a, k, l, expect)
}

func TestSample4(t *testing.T) {
	k, l := 2, 1
	a := []int{1, 2, 3, 4, 5, 6}
	expect := 0
	runSample(t, a, k, l, expect)
}
