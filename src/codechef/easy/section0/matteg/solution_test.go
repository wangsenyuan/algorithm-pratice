package main

import "testing"

func runSample(t *testing.T, R, C int, board [][]int, SX, SY int, N int, X []int, Y []int, expect int64) {
	res := solve(R, C, board, SX, SY, N, X, Y)
	if res != expect {
		t.Errorf("sample %d %d %v %d %d %d %v %v, expect %d, but got %d", R, C, board, SX, SY, N, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C, N := 5, 5, 2
	SX, SY := 2, 2
	X := []int{1, 2}
	Y := []int{2, 1}
	board := [][]int{
		{10, 11, 62, 14, 15},
		{57, 23, 34, 75, 21},
		{17, 12, 14, 11, 53},
		{84, 61, 24, 85, 22},
		{43, 89, 14, 15, 43},
	}
	expect := int64(188)
	runSample(t, R, C, board, SX, SY, N, X, Y, expect)
}

func TestSample2(t *testing.T) {
	R, C, N := 3, 3, 2
	SX, SY := 0, 0
	X := []int{1, 1}
	Y := []int{1, 1}
	board := [][]int{
		{9, 8, 7},
		{5, 6, 4},
		{1, 3, 2},
	}
	expect := int64(24)
	runSample(t, R, C, board, SX, SY, N, X, Y, expect)
}

func TestSample3(t *testing.T) {
	R, C, N := 2, 2, 1
	SX, SY := 1, 1
	X := []int{2}
	Y := []int{2}
	board := [][]int{
		{5, 6},
		{8, 3},
	}
	expect := int64(3)
	runSample(t, R, C, board, SX, SY, N, X, Y, expect)
}
