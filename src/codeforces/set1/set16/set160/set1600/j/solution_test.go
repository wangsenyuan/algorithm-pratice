package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat [][]int, expect []int) {
	res := solve(mat)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{9, 14, 11, 12, 13},
		{5, 15, 11, 6, 7},
		{5, 9, 14, 9, 14},
		{3, 2, 14, 3, 14},
	}
	expect := []int{9, 4, 4, 2, 1}
	runSample(t, mat, expect)
}
