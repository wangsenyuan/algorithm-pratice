package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P [][]int, expect bool) {
	res := solve(P)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	// check
	n := len(P)
	m := len(P[0])
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		Q[i] = make([]int, m)
	}

	for _, cur := range res {
		i, j, k := cur[0], cur[1], cur[2]
		Q[i-1][j-1] = k
		Q[i-1][j] = k
		Q[i][j-1] = k
		Q[i][j] = k
	}

	if !reflect.DeepEqual(P, Q) {
		t.Fatalf("Sample result %v, giving wrong result %v", res, Q)
	}
}

func TestSample1(t *testing.T) {
	P := [][]int{
		{5, 5, 3, 3},
		{1, 1, 5, 3},
		{2, 2, 5, 4},
		{2, 2, 4, 4},
	}
	expect := true
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := [][]int{
		{1, 1, 1, 1},
		{2, 2, 3, 1},
		{2, 2, 1, 1},
	}
	expect := false
	runSample(t, P, expect)
}
