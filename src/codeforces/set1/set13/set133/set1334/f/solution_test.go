package main

import "testing"

type Result struct {
	ok  bool
	ans int
}

func runSample(t *testing.T, a []int, cost []int, b []int, expect Result) {
	ok, ans := solve(a, cost, b)

	res := Result{ok, ans}

	if expect != res {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 1, 3, 3, 7, 8, 7, 9, 10, 7, 11}
	cost := []int{3, 5, 0, -2, 5, 3, 6, 7, 8, 2, 4}
	b := []int{3, 7, 10}
	expect := Result{true, 20}
	runSample(t, a, cost, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 5, 3, 6, 5}
	cost := []int{3, -9, 0, 16, 22, -14}
	b := []int{2, 3, 5, 6}
	expect := Result{false, 0}
	runSample(t, a, cost, b, expect)
}
