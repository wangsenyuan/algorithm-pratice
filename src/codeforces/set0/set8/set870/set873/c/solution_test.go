package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat [][]int, k int, expect []int) {
	res := solve(mat, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	k := 2
	expect := []int{4, 1}
	runSample(t, mat, k, expect)
}
