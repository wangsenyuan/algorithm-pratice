package main

import "testing"

func runSample(t *testing.T, m int, x int, A []int, H [][]int, expect int) {
	res := int(solve(m, x, A, H))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 4
	x := 3
	A := []int{1}
	H := [][]int{
		{5, 10, -6, -2},
	}
	expect := 5
	runSample(t, m, x, A, H, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	x := 1
	A := []int{1, 2, 1}
	H := [][]int{
		{1, 1},
		{2, 1},
		{1, 2},
	}
	expect := 5
	runSample(t, m, x, A, H, expect)
}

func TestSample3(t *testing.T) {
	m := 2
	x := 1
	A := []int{1, 2, 1}
	H := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
	}
	expect := 5
	runSample(t, m, x, A, H, expect)
}
