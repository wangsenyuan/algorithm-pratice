package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 1, 3, 2, 1}
	Q := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}
	expect := []int{7, 3, 3, 1}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 2, 1, 1073741823}
	Q := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
		{3, 4},
	}
	expect := []int{2, 3, 1, 1073741823}
	runSample(t, A, Q, expect)
}
