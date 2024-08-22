package main

import "testing"

func runSample(t *testing.T, n int, a [][]int, bans [][]int, expect []int) {
	res := solve(n, a, bans)

	getSum := func(arr []int) int {
		var res int
		for i := 0; i < n; i++ {
			res += a[i][arr[i]-1]
		}
		return res
	}

	x := getSum(expect)
	y := getSum(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := [][]int{
		{1, 2, 3},
		{1, 5},
		{2, 4, 6},
	}
	b := [][]int{
		{3, 2, 3},
		{3, 2, 2},
	}
	expect := []int{2, 2, 3}
	runSample(t, n, a, b, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	a := [][]int{
		{1, 2, 3},
		{1, 5},
		{2, 4, 6},
	}
	b := [][]int{
		{3, 2, 3},
		{2, 2, 3},
	}
	expect := []int{1, 2, 3}
	runSample(t, n, a, b, expect)
}
