package main

import (
	"testing"
)

func runSample(t *testing.T, n int, roads [][]int, a []int, expect bool) {
	res := solve(n, roads, a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	roads := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 6},
		{5, 6},
	}
	a := []int{0, 1, 2, 3, 0, 1}
	expect := true
	runSample(t, n, roads, a, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{3, 5},
	}
	a := []int{0, 1, 3, 2, 0}
	expect := false
	runSample(t, n, roads, a, expect)
}
