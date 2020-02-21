package main

import "testing"

func runSample(t *testing.T, N, M, W int, piles [][]int, disks [][]int, expect int64) {
	res := solve(N, M, W, piles, disks)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, M, W, piles, disks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, W := 11, 4, 13
	piles := [][]int{
		{19, 10},
		{8, 7},
		{11, 4},
		{26, 1},
		{4, 2},
		{15, 4},
		{19, 4},
		{1, 9},
		{4, 6},
		{19, 5},
		{15, 10},
	}
	disks := [][]int{
		{2, 1},
		{3, 100},
		{4, 10000},
		{5, 1000000},
	}
	expect := int64(206)
	runSample(t, N, M, W, piles, disks, expect)
}

func TestSample2(t *testing.T) {
	N, M, W := 11, 4, 13
	piles := [][]int{
		{19, 10},
		{8, 7},
		{11, 4},
		{26, 1},
		{4, 2},
		{15, 4},
		{19, 4},
		{1, 9},
		{4, 6},
		{19, 5},
		{15, 10},
	}
	disks := [][]int{
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 4},
	}
	expect := int64(5)
	runSample(t, N, M, W, piles, disks, expect)
}

func TestSample3(t *testing.T) {
	N, M, W := 1, 1, 1000000000
	piles := [][]int{
		{0, 500000000},
	}
	disks := [][]int{
		{1, 1},
	}
	expect := int64(-1)
	runSample(t, N, M, W, piles, disks, expect)
}
