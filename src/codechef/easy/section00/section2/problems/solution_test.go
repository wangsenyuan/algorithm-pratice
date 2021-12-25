package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P, S int, tasks [][]int, expect []int) {
	res := solve(P, S, tasks)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", P, S, tasks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	P, S := 3, 3
	tasks := [][]int{
		{16, 24, 60},
		{498, 861, 589},
		{14, 24, 62},
		{72, 557, 819},
		{16, 15, 69},
		{435, 779, 232},
	}
	expect := []int{2, 1, 3}
	runSample(t, P, S, tasks, expect)
}
