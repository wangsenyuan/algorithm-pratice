package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, tasks [][]int, expect []int64) {
	res := solve(n, tasks)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	tasks := [][]int{
		{5, 1},
		{3, 2},
		{5, 2},
		{1, 2},
		{4, 3},
		{4, 3},
		{4, 3},
		{5, 3},
		{1, 5},
		{3, 5},
		{1, 8},
		{2, 10},
	}
	expect := []int64{1, 3, 6, 8, 11, 17, 26, 32, 39, 46, 61, 71}
	runSample(t, n, tasks, expect)
}
