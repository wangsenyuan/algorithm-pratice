package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M, Q int, challeges [][]int, combos [][]int, expect []int64) {
	res := solve(N, M, Q, challeges, combos)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v, expect %v, but got %v", N, M, Q, challeges, combos, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, Q := 5, 3, 4
	challeges := [][]int{
		{1, 1, 3},
		{1, 2, -4},
		{4, 5, 2},
	}
	combos := [][]int{
		{1, 2},
		{1, 3},
		{1, 1},
		{2, 3},
	}

	expect := []int64{3, 0, 0, 4, 4}

	runSample(t, N, M, Q, challeges, combos, expect)
}
