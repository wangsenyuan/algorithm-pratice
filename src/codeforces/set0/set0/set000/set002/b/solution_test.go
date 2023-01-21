package main

import "testing"

func runSample(t *testing.T, n int, M [][]int, expect int) {
	res, path := solve(n, M)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
		return
	}


	x, y := 0, 0
	prod := M[0][0]

	for i := 0; i < len(path); i++ {
		if path[i] == 'D' {
			x++
		} else {
			y++
		}
		prod *= M[x][y]
	}

	if count(prod, 10) != expect {
		t.Errorf("Sample result %s, not correct",path)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	M := [][]int{
		{1, 2 ,3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := 0
	runSample(t, n, M, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	M := [][]int{
		{1, 2 ,3},
		{4, 0, 5},
		{5, 7, 9},
	}
	expect := 1
	runSample(t, n, M, expect)
}
