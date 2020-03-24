package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, A [][]int, expect []int) {
	_, res := solve(n, k, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	A := [][]int{
		{1, 2},
		{2, 1},
	}
	expect := []int{0, 0}
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 1
	A := [][]int{
		{1, 2},
	}
	expect := []int{2, 0}
	runSample(t, n, k, A, expect)
}
