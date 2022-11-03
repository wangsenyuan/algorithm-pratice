package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, p int, E [][]int, expect []int) {
	res := solve(n, p, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p := 4, 998244353
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
	}

	expect := []int{1, 6, 6}
	runSample(t, n, p, E, expect)
}

func TestSample2(t *testing.T) {
	n, p := 7, 100000007
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}

	expect := []int{1, 47, 340, 854, 880, 320}
	runSample(t, n, p, E, expect)
}

func TestSample3(t *testing.T) {
	n, p := 8, 1000000007
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
	}

	expect := []int{1, 126, 1806, 8400, 16800, 15120, 5040}
	runSample(t, n, p, E, expect)
}

func TestSample4(t *testing.T) {
	n, p := 2, 100000193
	E := [][]int{
		{2, 1},
	}

	expect := []int{1}
	runSample(t, n, p, E, expect)
}
