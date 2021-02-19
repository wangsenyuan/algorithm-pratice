package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, Q [][]int, expect []int) {
	res := solve(n, P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	P := []int{1, 1, 2, 2}
	Q := [][]int{
		{1, 10},
		{1, 20},
	}
	expect := []int{5, 0}
	runSample(t, n, P, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	P := []int{1, 2, 2, 3, 3}
	Q := [][]int{
		{1, 10},
		{1, 20},
	}
	expect := []int{5, 0}
	runSample(t, n, P, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	P := []int{1, 2, 3, 3, 4, 4}
	Q := [][]int{
		{2, 10},
		{3, 20},
	}
	expect := []int{5, 0}
	runSample(t, n, P, Q, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	P := []int{1, 2, 3, 3, 4, 4, 5}
	Q := [][]int{
		{2, 10},
		{3, 20},
		{5, 2},
	}
	expect := []int{5, 0, 0}
	runSample(t, n, P, Q, expect)
}
