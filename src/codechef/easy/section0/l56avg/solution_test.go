package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, A []int, expect float64) {
	res := solve(n, edges, A)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %v %v, expect %.6f, but got %.6f", n, edges, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{{1, 2}}
	A := []int{2, 3}
	expect := float64(2.5)
	runSample(t, n, edges, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{{5, 3}, {1, 3}, {2, 1}, {3, 4}}
	A := []int{4, 3, 5, 2, 1}
	expect := float64(2.6666667)
	runSample(t, n, edges, A, expect)
}
