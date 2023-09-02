package main

import "testing"

func runSample(t *testing.T, n int, k int, r int, c int, expect int) {
	res := solve(n, k, r, c)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
			if res[i][j] == 'X' {
				cnt[i][j]++
			}
		}
	}

	if expect != cnt[n-1][n-1] || res[r-1][c-1] != 'X' {
		t.Fatalf("Sample result %v, not correct", res)
	}

	get := func(a, b, c, d int) int {
		tmp := cnt[c][d]
		if a > 0 {
			tmp -= cnt[a][d]
		}
		if b > 0 {
			tmp -= cnt[c][b]
		}
		if a > 0 && b > 0 {
			tmp += cnt[a][b]
		}
		return tmp
	}

	for i := k - 1; i < n; i++ {
		for j := k - 1; j < n; j++ {
			if get(i-k, j-1, i, j) == 0 || get(i-1, j-k, i, j) == 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n, k, r, c := 3, 3, 3, 2
	expect := 3
	runSample(t, n, k, r, c, expect)
}

func TestSample2(t *testing.T) {
	n, k, r, c := 2, 1, 1, 2
	expect := 4
	runSample(t, n, k, r, c, expect)
}

func TestSample3(t *testing.T) {
	n, k, r, c := 6, 3, 4, 2
	expect := 12
	runSample(t, n, k, r, c, expect)
}
