package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, P []int, Q [][]int, expect []int) {
	res := solve(n, A, P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	P := []int{2, 3, 1}
	Q := [][]int{
		{1},
		{2, 2, 3},
		{3, 1},
	}
	expect := []int{3}
	runSample(t, n, A, P, Q, expect)
}
