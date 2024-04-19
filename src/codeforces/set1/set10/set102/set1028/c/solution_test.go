package main

import "testing"

func runSample(t *testing.T, rects [][]int, expect bool) {
	res := solve(rects)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var cnt int
	for _, rect := range rects {
		if cover(rect, res) {
			cnt++
		}
	}

	if cnt < len(rects)-1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func cover(rect []int, point []int) bool {
	x, y := point[0], point[1]
	return x >= rect[0] && x <= rect[2] && y >= rect[1] && y <= rect[3]
}

func TestSample1(t *testing.T) {
	rects := [][]int{
		{0, 0, 1, 1},
		{1, 1, 2, 2},
		{3, 0, 4, 1},
	}
	expect := true
	runSample(t, rects, expect)
}

func TestSample2(t *testing.T) {
	rects := [][]int{
		{0, 0, 5, 5},
		{0, 0, 4, 4},
		{1, 1, 4, 4},
		{1, 1, 4, 4},
	}
	expect := true
	runSample(t, rects, expect)
}

func TestSample3(t *testing.T) {
	rects := [][]int{
		{0, 0, 10, 8},
		{1, 2, 6, 7},
		{2, 3, 5, 6},
		{3, 4, 4, 5},
		{8, 1, 9, 2},
	}
	expect := true
	runSample(t, rects, expect)
}

func TestSample4(t *testing.T) {
	rects := [][]int{
		{-44, -47, -17, 5},
		{-40, -21, 46, 22},
		{-45, -47, -4, 7},
		{-45, -37, 17, 10},
		{-44, -24, 42, 6},
		{-40, -17, 22, 15},
		{-50, -46, -29, 17},
		{-19, -41, 9, 34},
		{-46, -24, 19, 27},
		{-42, -19, 20, 30},
	}
	expect := true
	runSample(t, rects, expect)
}
