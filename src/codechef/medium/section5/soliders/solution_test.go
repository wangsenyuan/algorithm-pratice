package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, p int, Q [][]int, expect []int) {
	res := solve(n, p, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	p := 0
	Q := [][]int{
		{3, 4},
		{2, 5},
		{7, 7},
	}
	expect := []int{1, 2, 3, 4, 5, 1, 7}
	runSample(t, n, p, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	p := 1
	Q := [][]int{
		{3, 4},
		{2, 5},
		{7, 7},
	}
	expect := []int{7, 5, 4, 3, 2, 7, 7}
	runSample(t, n, p, Q, expect)
}
