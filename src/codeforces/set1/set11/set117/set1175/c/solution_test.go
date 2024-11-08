package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 5}
	k := 2
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1000000000}
	k := 1
	expect := 500000000
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4}
	k := 0
	expect := 4
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 6}
	k := 0
	expect := 1
	runSample(t, a, k, expect)
}
