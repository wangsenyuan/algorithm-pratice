package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A [][]int, B [][]int, expect []int) {
	res := solve(n, A, B)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := [][]int{
		{1, 3},
		{1, 2},
		{3, 5},
		{3, 4},
	}
	B := [][]int{
		{1, 5},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := []int{0, 1, 1, 2, 1}
	runSample(t, n, A, B, expect)
}
