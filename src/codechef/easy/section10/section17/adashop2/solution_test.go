package main

import "testing"

func runSample(t *testing.T, r0, c0 int) {
	res := solve(r0, c0)

	if len(res) > 64 {
		t.Fatalf("Sample %d %d, result %v, has more than 64 steps", r0, c0, res)
	}

	board := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		board[i] = make([]bool, 9)
	}

	for _, step := range res {
		u, v := step[0], step[1]

		if u%2 != u%2 || u < 0 || u > 8 || v < 0 || v > 8 {
			t.Fatalf("steps %v, not a valid black cell", step)
		}

		board[u][v] = true
	}
	var cnt int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] {
				cnt++
			}
		}
	}
	if cnt != 32 {
		t.Fatalf("cover only %d black cells", cnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3)
}
