package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, a [][]int, expect []int) {
	res := solve(n, k, a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 1, 2
	a := [][]int{
		{-5},
	}
	expect := []int{0, -5}
	runSample(t, n, k, a, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 4
	a := [][]int{
		{2, -3},
		{-1},
	}
	expect := []int{2, 0, -1, -3}
	runSample(t, n, k, a, expect)
}
