package main

import "testing"

func runSample(t *testing.T, n, m int, expect int) {
	k, res := solve(n, m)
	if k != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, k)
	} else if !valid(n, m, res) {
		t.Errorf("Sample %d %d, result %v, not valid", n, m, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func valid(n, m int, mat [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ns := make(map[int]bool)
			var cnt int
			for k := 0; k < 4; k++ {
				x, y := i+dd[k], j+dd[k+1]
				if x >= 0 && x < n && y >= 0 && y < m {
					cnt++
					ns[mat[x][y]] = true
				}
			}
			if len(ns) != cnt {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 3)
}
