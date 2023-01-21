package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, segs [][]int, expect []int) {
	res := solve(len(segs), segs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{2, 4, 20},
		{7, 8, 22},
	}
	expect := []int{20, 42}
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{5, 11, 42},
		{5, 11, 42},
	}
	expect := []int{42, 42}
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 4, 4},
		{5, 8, 9},
		{7, 8, 7},
		{2, 10, 252},
		{1, 11, 271},
		{1, 10, 1},
	}
	expect := []int{4, 13, 11, 256, 271, 271}
	runSample(t, segs, expect)
}
