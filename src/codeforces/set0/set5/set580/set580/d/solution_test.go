package main

import "testing"

func runSample(t *testing.T, m int, a []int, restrictions [][]int, expect int) {
	res := solve(m, a, restrictions)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	a := []int{1, 1}
	restrictions := [][]int{
		{2, 1, 1},
	}
	expect := 3
	runSample(t, m, a, restrictions, expect)
}

func TestSample2(t *testing.T) {
	m := 3
	a := []int{1, 2, 3, 4}
	restrictions := [][]int{
		{2, 1, 5},
		{3, 4, 2},
	}
	expect := 12
	runSample(t, m, a, restrictions, expect)
}
