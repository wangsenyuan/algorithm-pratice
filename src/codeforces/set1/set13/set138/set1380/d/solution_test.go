package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, cost1 int, cost2 int, expect int) {
	res := solve(len(a), a, b, k, cost1, cost2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	cost1 := 5
	cost2 := 3
	a := []int{3, 1, 4, 5, 2}
	b := []int{3, 5}
	expect := 8
	runSample(t, a, b, k, cost1, cost2, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	cost1 := 5
	cost2 := 4
	a := []int{4, 3, 1, 2}
	b := []int{2, 4, 3, 1}
	expect := -1
	runSample(t, a, b, k, cost1, cost2, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	cost1 := 2
	cost2 := 11
	a := []int{1, 3, 2, 4}
	b := []int{1, 3, 2, 4}
	expect := 0
	runSample(t, a, b, k, cost1, cost2, expect)
}
