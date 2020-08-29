package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	_, b, ok := countNeighbours(res)

	if !ok || b != n {
		t.Errorf("Sample %d, result %v expect even neighbours, but got not", n, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func countNeighbours(coords [][]int) (int, int, bool) {
	mem := make(map[[2]int]int)

	for _, coord := range coords {
		mem[[...]int{coord[0], coord[1]}]++
	}

	var a, b int

	for _, coord := range coords {
		x, y := coord[0], coord[1]

		var c int

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			c += mem[[...]int{u, v}]
		}

		if c == 2 {
			a++
		} else if c == 4 {
			b++
		} else {
			return -1, -1, false
		}
	}

	return a, b, true
}

func TestSample1(t *testing.T) {
	runSample(t, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 5)
}

func TestSample3(t *testing.T) {
	for i := 1; i <= 100; i++ {
		runSample(t, i)
	}
}
