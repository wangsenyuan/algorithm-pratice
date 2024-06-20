package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, p []int, x int) {
	n := len(p)
	a := make([]int, n)
	copy(a, p)
	res := solve(a, x)

	if len(res) > 2 {
		t.Fatalf("Sample result %v took too much steps (>2)", res)
	}

	for _, step := range res {
		i, j := step[0]-1, step[1]-1
		p[i], p[j] = p[j], p[i]
	}

	l, r := 1, n+1

	for r-l > 1 {
		mid := (l + r) / 2
		if p[mid-1] <= x {
			l = mid
		} else {
			r = mid
		}
	}

	if p[l-1] != x {
		t.Fatalf("Sample result %v, not find expect %d", res, x)
	}
}

func TestSample1(t *testing.T) {
	p := []int{6, 5, 4, 3, 2, 1}
	x := 3
	runSample(t, p, x)
}

func TestRandomSample(t *testing.T) {
	n := 10
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}

	for x := 0; x < 10; x++ {
		rand.Shuffle(n, func(i, j int) {
			p[i], p[j] = p[j], p[i]
		})
		runSample(t, p, x+1)
	}
}
