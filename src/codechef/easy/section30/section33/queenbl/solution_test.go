package main

import "testing"

func runSample(t *testing.T, king []int, expect int) {
	res := solve(king)

	if res[king[0]-1][king[1]-1] != 1 {
		t.Fatalf("grid should have king at %v", king)
	}

	king[0]--
	king[1]--

	grid := make([][]int, 8)
	for i := 0; i < 8; i++ {
		grid[i] = make([]int, 8)
	}

	var cnt int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if res[i][j] == 2 {
				cnt++
				// a queen
				if i == king[0] || j == king[1] || abs(king[0]-i) == abs(king[1]-j) {
					t.Fatalf("queen can attack king at %d %d", i+1, j+1)
				}
				for x := i - 1; x >= 0; x-- {
					grid[x][j] = 2
				}
				for x := i + 1; x < 8; x++ {
					grid[x][j] = 2
				}
				for y := j - 1; y >= 0; y-- {
					grid[i][y] = 2
				}
				for y := j + 1; y < 8; y++ {
					grid[i][y] = 2
				}
				for x := 1; i-x >= 0 && j-x >= 0; x++ {
					grid[i-x][j-x] = 2
				}
				for x := 1; i+x < 8 && j+x < 8; x++ {
					grid[i+x][j+x] = 2
				}
				for x := 1; i-x >= 0 && j+x < 8; x++ {
					grid[i-x][j+x] = 2
				}
				for x := 1; i+x < 8 && j-x >= 0; x++ {
					grid[i+x][j-x] = 2
				}
			}
		}
	}

	if cnt != expect {
		t.Fatalf("Sample %v expect %d, but got %d with %v", king, expect, cnt, res)
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			u, v := king[0]+i, king[1]+j
			if u >= 0 && u < 8 && v >= 0 && v < 8 {
				if grid[u][v] != 2 {
					t.Fatalf("sample result %v, can't win against king %v", res, king)
				}
			}
		}
	}
}

func TestSample1(t *testing.T) {
	king := []int{2, 1}
	expect := 2
	runSample(t, king, expect)
}

func TestSample(t *testing.T) {
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			king := []int{i, j}
			expect := 1
			if i > 1 && i < 8 || j > 1 && j < 8 {
				expect++
			}
			runSample(t, king, expect)
		}
	}
}
