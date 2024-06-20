package main

import "testing"

func runSample(t *testing.T, segs [][]int, expect bool) {
	res := solve(segs)

	if res[0] > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	a := segs[res[0]-1]
	b := segs[res[1]-1]
	if a[0] < b[0] || a[1] > b[1] {
		t.Fatalf("Sample result %v, (%v %v), not correct", res, a, b)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{1, 10},
		{2, 9},
		{3, 9},
		{2, 3},
		{2, 9},
	}
	expect := true
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{1, 5},
		{2, 6},
		{6, 20},
	}
	expect := false
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 10},
		{10, 20},
		{9, 11},
	}
	expect := false
	runSample(t, segs, expect)
}
