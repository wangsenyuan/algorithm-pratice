package main

import "testing"

func runSample(t *testing.T, W []int, friends [][]int, can bool) {
	res := solve(W, friends)

	if can != (len(res) > 0) {
		t.Errorf("sample %v %v, expect %t, but got %v", W, friends, can, res)
		return
	}
	if !can {
		return
	}
	for _, cur := range res {
		x, y := friends[cur-1][0], friends[cur-1][1]
		if W[x] == 0 && W[y] == 0 {
			t.Fatalf("Sample res %v, not correct", res)
		}
		W[x] = max(0, W[x]-1)
		W[y] = max(0, W[y]-1)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	W := []int{1, 2, 1}
	friends := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	runSample(t, W, friends, true)
}

func TestSample2(t *testing.T) {
	W := []int{1, 1, 0}
	friends := [][]int{
		{1, 2},
		{1, 3},
	}
	runSample(t, W, friends, true)
}

func TestSample3(t *testing.T) {
	W := []int{1, 2, 0, 1}
	friends := [][]int{
		{1, 3},
		{1, 2},
		{2, 3},
		{2, 4},
	}
	runSample(t, W, friends, true)
}

func TestSample4(t *testing.T) {
	W := []int{1, 1, 1, 2, 1}
	friends := [][]int{
		{3, 4},
		{1, 2},
		{2, 3},
		{4, 5},
		{4, 5},
	}
	runSample(t, W, friends, true)
}

func TestSample5(t *testing.T) {
	W := []int{2, 4, 1, 4}
	friends := [][]int{
		{3, 2},
		{4, 2},
		{4, 1},
		{3, 1},
		{4, 1},
		{1, 3},
		{3, 2},
		{2, 1},
		{3, 1},
		{2, 4},
	}
	runSample(t, W, friends, false)
}
