package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []float64) {
	res := solve(n, edges)

	if len(expect) == 0 {
		if len(res) != 0 {
			t.Fatalf("Sample expect no result, but got %v", res)
		}
		return
	}

	x := getSum(expect)
	y := getSum(res)

	if math.Abs(x-y)/math.Max(1.0, y) > 1e-8 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		u--
		v--
		wf := float64(w)
		rf := res[u] + res[v]
		if math.Abs(rf-wf)/wf > 1e-8 {
			t.Fatalf("Sample result %v, have a wrong assigment for edge %v, got %f", res, edge, rf)
		}
	}
}

func getSum(arr []float64) float64 {
	var res float64
	for _, x := range arr {
		res += x
	}
	return res
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 2},
		{3, 4, 1},
	}
	expect := []float64{0.5, 0.5, 1.5, -0.5}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2, 1},
	}
	expect := []float64{0.3, 0.7}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 2},
		{2, 3, 2},
	}
	expect := []float64{0, 2, 0}
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 2},
		{2, 2, 1},
		{2, 1, 1},
		{1, 2, 2},
	}
	runSample(t, n, edges, nil)
}
