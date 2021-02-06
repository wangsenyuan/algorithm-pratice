package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int, A []int, Q [][]int, expect []int) {
	res := solve(n, m, k, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 3, 7, 2
	A := []int{1, 2, 4}
	Q := [][]int{
		{1, 3},
		{2, 3},
	}
	expect := []int{3, -1}
	runSample(t, n, m, k, A, Q, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 3, 3, 3
	A := []int{1, 2, 3}
	Q := [][]int{
		{1, 2},
		{1, 3},
		{2, 2},
	}
	expect := []int{2, 1, -1}
	runSample(t, n, m, k, A, Q, expect)
}
