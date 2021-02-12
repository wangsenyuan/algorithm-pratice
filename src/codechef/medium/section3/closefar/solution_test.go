package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, E [][]int, m int, Q [][]int, expect []int64) {
	res := solve(n, A, E, m, Q)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 7, 4, 5}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}
	m := 7
	Q := [][]int{
		{X_C, 1, 5},
		{X_F, 1, 5},
		{X_C, 2, 4},
		{X_C, 1, 2},
		{X_F, 1, 3},
		{X_F, 3, 4},
		{X_F, 2, 4},
	}
	expect := []int64{1, 4, 2, 1, 6, 5, 2}

	runSample(t, n, A, E, m, Q, expect)
}
