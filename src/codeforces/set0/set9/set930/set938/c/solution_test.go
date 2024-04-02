package main

import "testing"

func runSample(t *testing.T, x int, expect []int) {
	res := solve(x)
	if (len(expect) > 0) != (len(res) > 0) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if len(res) == 0 {
		return
	}

	n, m := res[0], res[1]
	l := n / m
	if n*n-l*l != x {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 21, []int{5, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, 0, []int{1, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, 1, nil)
}
