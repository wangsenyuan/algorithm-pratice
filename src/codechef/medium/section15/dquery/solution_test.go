package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int64) {
	res := solve1(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{2, 3},
		{3, 4},
	}
	expect := []int64{8, 10}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 4, 1, 8, 9, 2, 4, 10, 7, 20}
	Q := [][]int{
		{2, 5},
		{3, 8},
		{5, 6},
		{2, 3},
		{7, 10},
	}
	expect := []int64{43, 41, 27, 24, 68}
	runSample(t, A, Q, expect)
}
