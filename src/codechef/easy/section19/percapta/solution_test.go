package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, B []int, E [][]int, expect []int) {
	res := solve(n, A, B, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{10, 1, 2}
	B := []int{5, 1, 1}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := []int{1, 3}
	runSample(t, n, A, B, E, expect)
}
