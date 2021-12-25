package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect [][]int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, bug got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 4, 1, 5}
	expect := [][]int{
		{2, 4},
		{2, 3},
		{2, 3},
		{-1, -1},
	}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 2, 1}
	expect := [][]int{
		{2, 2},
		{-1, -1},
	}
	runSample(t, A, expect)
}
