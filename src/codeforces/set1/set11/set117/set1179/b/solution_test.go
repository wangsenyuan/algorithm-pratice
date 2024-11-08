package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, m int) {
	res := solve(n, m)

	if len(res) != n*m {
		t.Fatalf("Sample result %v, not correct", res)
	}

	grid := make([][]bool, n+1)
	for i := 1; i <= n; i++ {
		grid[i] = make([]bool, m+1)
	}

	type pair struct {
		first  int
		second int
	}

	freq := make(map[pair]int)

	for i, cur := range res {
		x, y := cur[0], cur[1]
		if grid[x][y] {
			t.Fatalf("Sample result %v, not correct, visit same cell twice", res)
		}
		if i > 0 {
			dx := x - res[i-1][0]
			dy := y - res[i-1][1]
			p := pair{dx, dy}
			freq[p]++
			if freq[p] > 1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample3(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(100) + 1
		m := rand.Intn(100) + 1
		runSample(t, n, m)
	}
}
