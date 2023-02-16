package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, Q [][]int, expect []int) {
	res := solve(n, m, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v,  but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 20
	Q := [][]int{
		{1, 300000, 2},
		{2, 400000, 2},
		{1, 1000000, 3},
	}
	expect := []int{-1, -1, 1, -1, -1, 1, -1, -1, -1, 3, -1, 2, 3, -1, -1, 3, -1, -1, -1, 3}
	runSample(t, n, m, Q, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 20
	Q := [][]int{
		{1, 399999, 2},
		{2, 412345, 2},
		{1, 1000001, 3},
	}
	expect := []int{-1, -1, -1, 1, -1, -1, -1, 1, -1, -1, 3, -1, -1, -1, 3, -1, 2, -1, 3, -1}
	runSample(t, n, m, Q, expect)
}
