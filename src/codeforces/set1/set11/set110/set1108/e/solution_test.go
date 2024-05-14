package main

import "testing"

func runSample(t *testing.T, a []int, segs [][]int, expect int) {
	res, set := solve(a, segs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	n := len(a)

	diff := make([]int, n+1)
	for _, i := range set {
		l, r := segs[i-1][0]-1, segs[i-1][1]-1
		diff[l]--
		diff[r+1]++
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	x, y := oo, -oo
	for i := 0; i < n; i++ {
		diff[i] += a[i]
		x = min(x, diff[i])
		y = max(y, diff[i])
	}

	if y-x != expect {
		t.Fatalf("Sample result %v, get wrong answer %d", set, y-x)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, -2, 3, 1, 2}
	segs := [][]int{
		{1, 3},
		{4, 5},
		{2, 5},
		{1, 3},
	}
	expect := 6
	runSample(t, a, segs, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, -2, 3, 1, 4}
	segs := [][]int{
		{3, 5},
		{3, 4},
		{2, 4},
		{2, 5},
	}
	expect := 7
	runSample(t, a, segs, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1000000}

	expect := 0
	runSample(t, a, nil, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-4, 7}
	segs := [][]int{
		{2, 2},
		{2, 2},
	}
	expect := 11
	runSample(t, a, segs, expect)
}
