package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 6
	a := []int{1, 1, 2, 3, 4, 5, 6}
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{5, 1, 3, 2, 3}
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	a := []int{5, 5, 5, 5, 4}
	expect := -1
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	k := 4
	a := []int{1, 2, 3, 3, 2, 2, 5, 5}
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	k := 3
	a := []int{1, 2, 2, 3, 5, 3, 7, 2, 4, 3}
	expect := 0
	runSample(t, a, k, expect)
}
