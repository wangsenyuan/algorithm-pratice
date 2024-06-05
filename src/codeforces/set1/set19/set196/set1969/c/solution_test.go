package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := []int{3, 1, 2}
	expect := 4
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	a := []int{5}
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	a := []int{2, 2, 1, 3}
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	a := []int{4, 1, 2, 2, 4, 3}
	expect := 10
	runSample(t, a, k, expect)
}
