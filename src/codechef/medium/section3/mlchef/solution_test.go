package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, HP [][]int, Q int, input [][]int, expect []int) {
	H := make([]int, N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = HP[i][0]
		P[i] = HP[i][1]
	}
	res := solve(N, H, P, Q, input)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v expect %v, but got %v", N, HP, Q, input, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 4
	HP := [][]int{
		{1, 0},
		{2, 0},
		{2, 2},
		{1, 2},
	}
	Q := 4
	input := [][]int{
		{1, 2, 1},
		{2, 2},
		{1, 0, 1},
		{2, 0},
	}
	expect := []int{1, 1}
	runSample(t, N, HP, Q, input, expect)
}
