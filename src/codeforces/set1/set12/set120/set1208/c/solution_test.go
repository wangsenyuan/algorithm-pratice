package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	row := make([]int, n)
	col := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row[i] ^= res[i][j]
			col[j] ^= res[i][j]
		}
	}

	x := row[0]
	for i := 0; i < n; i++ {
		if x != row[i] || x != col[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 12)
}

func TestSample4(t *testing.T) {
	runSample(t, 100)
}
