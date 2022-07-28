package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 6, 9}
	Q := [][]int{
		{1, 1},
		{2, 2},
		{2, 3},
	}
	expect := []int{3, 1, 2}
	runSample(t, A, Q, expect)
}
