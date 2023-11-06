package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, segs [][]int, queries [][]int, expect []int) {
	res := solve(k, segs, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	segs := [][]int{
		{91, 94},
		{92, 97},
		{97, 99},
	}
	queries := [][]int{
		{92, 94},
		{93, 97},
		{95, 96},
		{90, 100},
	}
	expect := []int{3, 3, 0, 4}
	runSample(t, k, segs, queries, expect)
}
