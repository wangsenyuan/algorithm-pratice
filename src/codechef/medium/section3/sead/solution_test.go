package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, m int, quries [][]int, expect []int) {
	res := solve(n, A, m, quries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, m, quries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 10, 50}
	m := 6
	queries := [][]int{
		{1, 1},
		{5, 3},
		{11, 7},
		{100000, 1},
		{1000000, 1000000},
		{11, 6},
	}
	expect := []int{1, 1, 1, 5, 1, 4}
	runSample(t, n, A, m, queries, expect)
}
