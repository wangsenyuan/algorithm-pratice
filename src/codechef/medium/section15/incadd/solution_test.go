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
	A := []int{3, 2, 1}
	Q := [][]int{
		{3, 10},
		{2, 1},
	}
	expect := []int{1, 2}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2, 2}
	Q := [][]int{
		{4, 2},
		{4, 1},
	}
	expect := []int{0, 1}
	runSample(t, A, Q, expect)
}
