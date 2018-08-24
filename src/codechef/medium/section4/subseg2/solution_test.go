package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, courses [][]int, plans [][]int, expect []int) {
	res := solve(N, Q, courses, plans)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, courses, plans, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 3, 3
	courses := [][]int{
		{1, 3},
		{5, 6},
		{2, 4},
	}
	plans := [][]int{
		{1, 6},
		{1, 3},
		{2, 3},
	}
	expect := []int{2, 1, 0}
	runSample(t, N, Q, courses, plans, expect)
}
