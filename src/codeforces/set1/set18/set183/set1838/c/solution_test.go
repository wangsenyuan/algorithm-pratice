package main

import "testing"

var dd = []int{-1, 0, 1, 0, -1}

func runSample(t *testing.T, n int, m int) {
	res := solve(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				u, v := i+dd[k], j+dd[k+1]
				if u >= 0 && u < n && v >= 0 && v < m {
					diff := abs(res[i][j] - res[u][v])
					if diff >= len(lpf) || lpf[diff] == diff {
						t.Fatalf("Sample result %v, not correct for %d %d", res, n, m)
					}
				}
			}
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 4)
}

func TestSample3(t *testing.T) {
	for n := 4; n <= 30; n++ {
		for m := 4; m <= 40; m++ {
			runSample(t, n, m)
		}
	}
}

func TestSample4(t *testing.T) {
	n, m := 13, 97
	runSample(t, n, m)
}
