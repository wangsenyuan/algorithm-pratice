package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int) {
	res := solve(n, m, k)
	score := make([]int, n)
	x := n / m
	y := (n + m - 1) / m
	for _, game := range res {
		for _, table := range game {
			for _, i := range table {
				if x != y && len(table) == y {
					score[i-1]++
				}
			}
		}
	}
	sort.Ints(score)
	if score[0] < score[n-1]-1 {
		t.Fatalf("Sample result %v, not giving correct answer", res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 5, 2, 2
	runSample(t, n, m, k)
}

func TestSample2(t *testing.T) {
	n, m, k := 8, 3, 1
	runSample(t, n, m, k)
}

func TestSample3(t *testing.T) {
	n, m, k := 2, 1, 3
	runSample(t, n, m, k)
}
