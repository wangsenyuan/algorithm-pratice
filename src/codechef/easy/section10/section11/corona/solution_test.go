package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M, K int, E [][]int, C [][]int, expect []int64) {
	res := solve(N, M, K, E, C)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, K := 3, 2, 2
	C := [][]int{
		{1, 5},
		{2, 10},
	}
	E := [][]int{
		{1, 2, 6},
		{2, 3, 9},
	}
	expect := []int64{5, 10, 19}
	runSample(t, N, M, K, E, C, expect)
}

func TestSample2(t *testing.T) {
	N, M, K := 3, 2, 2
	C := [][]int{
		{1, 5},
		{2, 10},
	}
	E := [][]int{
		{1, 2, 1},
		{2, 3, 9},
	}
	expect := []int64{5, 6, 15}
	runSample(t, N, M, K, E, C, expect)
}
