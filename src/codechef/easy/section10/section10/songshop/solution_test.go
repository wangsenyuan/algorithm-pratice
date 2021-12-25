package main

import "testing"

func runSample(t *testing.T, N, M, P int, songs [][]int, albums []int, expect int64) {
	res := solve(N, M, P, songs, albums)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, M, P, songs, albums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, P := 5, 2, 24
	songs := [][]int{
		{1, 7, 2},
		{1, 5, 2},
		{1, 4, 1},
		{2, 9, 1},
		{2, 13, 2},
	}
	albums := []int{10, 15}
	runSample(t, N, M, P, songs, albums, 7)
}
