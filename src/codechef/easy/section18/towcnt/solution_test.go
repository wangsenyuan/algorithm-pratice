package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, H int, A [][]int, expect []int) {
	res := solve(n, H, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, H := 4, 10
	A := [][]int{
		{0, 2, 5},
		{1, 5, 6},
		{1, 8, 8},
		{0, 11, 2},
	}
	expect := []int{2, 3, 2, 3}
	runSample(t, n, H, A, expect)
}

func TestSample2(t *testing.T) {
	n, H := 3, 10
	A := [][]int{
		{0, 2, 2},
		{0, 2, 3},
		{0, 2, 4},
	}
	expect := []int{1, 2, 1}
	runSample(t, n, H, A, expect)
}
