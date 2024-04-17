package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, A []int, B []int, pb []int, expect int) {
	pa := make([]int, n)
	for i := 0; i < n; i++ {
		pa[i] = i + 1
	}

	ask := func(x byte, y int) int {
		if x == 'A' {
			for i := 0; i < n; i++ {
				if pa[i] == y {
					return pb[i]
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if pb[i] == y {
					return pa[i]
				}
			}
		}

		return 0
	}

	res := solve(n, edges, A, B, ask)

	if (expect < 0) != (res < 0) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	if expect < 0 {
		return
	}
	var flag int
	for i := 0; i < n; i++ {
		if pa[i] == res {
			flag |= 1
		}
		if pb[i] == res {
			flag |= 2
		}
	}
	if flag != 3 {
		t.Fatalf("Sample result %d, not correct, it occurence flag is %d", res, flag)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	A := []int{1}
	B := []int{2}
	pb := []int{2, 3, 1}
	expect := 1
	runSample(t, n, edges, A, B, pb, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{4, 5},
		{4, 6},
	}
	A := []int{1, 3, 4, 5}
	B := []int{3, 5, 2}
	pb := []int{5, 3, 2, 4, 1, 6}
	expect := 1
	runSample(t, n, edges, A, B, pb, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{4, 5},
		{4, 6},
	}
	A := []int{1, 2, 3}
	B := []int{4, 1, 6}
	pb := []int{5, 3, 2, 4, 1, 6}
	expect := -1
	runSample(t, n, edges, A, B, pb, expect)
}
