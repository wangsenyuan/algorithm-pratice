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
	A := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{1, 5, 4},
		{2, 5, 2},
	}
	expect := []int{3, 2}
	runSample(t, A, Q, expect)
}
