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
	A := []int{1, 2, 3}
	Q := [][]int{
		{1, 4},
		{3, 0},
	}
	expect := []int{3, 7, 6}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	Q := [][]int{
		{4, 0},
	}
	expect := []int{7, 3}
	runSample(t, A, Q, expect)
}
