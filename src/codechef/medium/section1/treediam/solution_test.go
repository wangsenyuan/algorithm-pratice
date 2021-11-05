package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, W []int, E [][]int, S []int, expect []int) {
	res := solve(n, W, E, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 3
	W := []int{1, 2, 3}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	S := []int{2, 1}
	expect := []int{6, 9, 6}
	runSample(t, n, W, E, S, expect)
}
