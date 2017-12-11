package main

import "testing"

func TestSample(t *testing.T) {
	n := 10
	A := [][]byte{
		[]byte(string("..X....XX.")),
		[]byte(string("XX.X..XX.X")),
		[]byte(string(".....XX..X")),
		[]byte(string(".XXX..X.X.")),
		[]byte(string(".....X..XX")),
		[]byte(string("....X....X")),
		[]byte(string("X.X....XX.")),
		[]byte(string(".X...X.X.X")),
		[]byte(string("X.X..X....")),
		[]byte(string("..XXXXX.XX")),
	}

	B := [][]int{
		{0, 0, 2, 0, 0, 0, 0, 3, 3, 0},
		{2, 2, 0, 2, 0, 0, 3, 3, 0, 2},
		{0, 0, 0, 0, 0, 3, 3, 0, 0, 2},
		{0, 3, 3, 3, 0, 0, 3, 0, 2, 0},
		{0, 0, 0, 0, 0, 3, 0, 0, 2, 2},
		{0, 0, 0, 0, 3, 0, 0, 0, 0, 3},
		{4, 0, 3, 0, 0, 0, 0, 2, 3, 0},
		{0, 4, 0, 0, 0, 3, 0, 3, 0, 2},
		{3, 0, 4, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 5, 5, 5, 5, 5, 0, 2, 2},
	}

	res := solve(n, A)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] != B[i][j] {
				t.Errorf("answer at (%d, %d) is wrong: %d vs %d", i, j, res[i][j], B[i][j])
			}
		}
	}
}
