package main

import "testing"

func runSample(t *testing.T, mat [][]int, expect bool) {
	res, k := solve(mat)

	if expect != (k > 0) {
		t.Fatalf("Sample expect %t, but got %s %d", expect, res, k)
	}

	if !expect {
		return
	}
	const inf = 1 << 30

	tmp := []int{0, inf, inf, 0}

	for i := 0; i < len(mat); i++ {
		r := res[i] == 'R'

		for j := 0; j < len(mat[0]); j++ {
			if j < k {
				// left part
				if r {
					tmp[1] = min(tmp[1], mat[i][j])
				} else {
					tmp[0] = max(tmp[0], mat[i][j])
				}
			} else {
				if r {
					tmp[3] = max(tmp[3], mat[i][j])
				} else {
					tmp[2] = min(tmp[2], mat[i][j])
				}
			}
		}
	}

	if tmp[0] >= tmp[1] || tmp[2] <= tmp[3] {
		t.Fatalf("Sample result %s %d, not correct", res, k)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 5, 8, 8, 7},
		{5, 2, 1, 4, 3},
		{1, 6, 9, 7, 5},
		{9, 3, 3, 3, 2},
		{1, 7, 9, 9, 8},
	}
	expect := true
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{8, 9, 8},
		{1, 5, 3},
		{7, 5, 7},
	}
	expect := false
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{3, 3, 3, 2, 2, 2},
		{1, 1, 1, 4, 4, 4},
	}
	expect := true
	runSample(t, mat, expect)
}

func TestSample4(t *testing.T) {
	mat := [][]int{
		{1, 12, 7},
		{13, 14, 3},
	}
	expect := true
	runSample(t, mat, expect)
}
