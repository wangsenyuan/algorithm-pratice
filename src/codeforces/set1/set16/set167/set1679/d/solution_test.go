package main

import "testing"

func runSample(t *testing.T, n int, A []int, E [][]int, k int64, expect int) {
	res := solve(n, A, E, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	var k int64 = 4
	A := []int{1, 10, 2, 3, 4, 5}
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 2},
		{2, 5},
	}
	expect := 4
	runSample(t, n, A, E, k, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	var k int64 = 5
	A := []int{1, 1}
	E := [][]int{
		{1, 2},
	}
	expect := -1
	runSample(t, n, A, E, k, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	var k int64 = 1
	A := []int{1000000000}
	expect := 1000000000
	runSample(t, n, A, nil, k, expect)
}
