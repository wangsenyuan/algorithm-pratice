package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q [][]int) {
	res := solve(n, A, Q)
	expect := solve1(n, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{7, 4, 2, 9, 3}
	Q := [][]int{
		{7, 2, 4},
		{8, 1, 5},
	}
	runSample(t, n, A, Q)
}
