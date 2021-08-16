package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, V []int, R [][]int, expect []int) {
	res := solve(k, V, R)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	V := []int{3, 1, 2, 2, 1, 1, 2, 1}
	R := [][]int{
		{0, 5},
		{2, 4},
		{1, 7},
	}
	expect := []int{1, 0, 2}
	k := 2
	runSample(t, k, V, R, expect)
}
