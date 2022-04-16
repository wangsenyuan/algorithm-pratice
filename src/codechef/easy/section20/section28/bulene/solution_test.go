package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S []int, B [][]int, expect []int) {
	res := solve(S, B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []int{51, 84, 24, 19, 77, 10, 34}
	B := [][]int{
		{3, 2},
		{6, 2},
		{3, 6},
		{10, 2},
	}
	expect := []int{29, 62, 21, 16, 74, 7, 34}
	runSample(t, S, B, expect)
}

func TestSample2(t *testing.T) {
	S := []int{10, 5, 4, 8}
	B := [][]int{
		{4, 2},
		{2, 4},
		{3, 2},
	}
	expect := []int{1, 0, 0, 6}
	runSample(t, S, B, expect)
}
