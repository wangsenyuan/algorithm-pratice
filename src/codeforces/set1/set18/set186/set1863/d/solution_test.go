package main

import "testing"

func runSample(t *testing.T, n int, m int, grid []string, expect bool) {
	res := solve(n, m, grid)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	col := make([]int, m)

	for i := 0; i < n; i++ {
		var row int
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' {
				if res[i][j] != '.' {
					t.Fatalf("Sample result %v, not correct", res)
				}
				continue
			}
			if res[i][j] == 'W' {
				row++
				col[j]++
			} else {
				row--
				col[j]--
			}
			if grid[i][j] == 'L' {
				if res[i][j+1] == '.' || res[i][j+1] == res[i][j] {
					t.Fatalf("Sample result %v, not correct", res)
				}
			} else if grid[i][j] == 'U' {
				if res[i+1][j] == '.' || res[i+1][j] == res[i][j] {
					t.Fatalf("Sample result %v, not correct", res)
				}
			}
		}

		if row != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	for j := 0; j < m; j++ {
		if col[j] != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 6
	grid := []string{
		"..LR..",
		"ULRU..",
		"DLRDUU",
		"..LRDD",
	}
	expect := true
	runSample(t, n, m, grid, expect)
}
