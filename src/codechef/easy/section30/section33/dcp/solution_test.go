package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, Q []int, E [][]int, expect []int) {
	res := solve(n, m, Q, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 2
	Q := []int{5, 10, 12, 2}
	E := [][]int{
		{1, 2},
		{3, 2, 4, 3},
		{3, 1},
		{2, 4},
	}
	expect := []int{0, 25, 0, 66}
	runSample(t, n, m, Q, E, expect)
}
