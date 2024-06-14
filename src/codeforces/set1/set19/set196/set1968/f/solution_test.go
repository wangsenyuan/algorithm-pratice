package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []bool) {
	res := solve(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 3, 0}
	queries := [][]int{
		{1, 5},
		{2, 4},
		{3, 5},
		{1, 3},
		{3, 4},
	}
	expect := []bool{true, true, false, false, false}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	queries := [][]int{
		{1, 5},
		{2, 4},
		{3, 5},
		{1, 3},
		{2, 3},
	}
	expect := []bool{true, false, false, true, false}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{12, 9, 10, 9, 10, 11, 9}
	queries := [][]int{
		{1, 5},
		{1, 7},
		{2, 6},
		{2, 7},
	}
	expect := []bool{false, false, false, false}
	runSample(t, nums, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1}
	queries := [][]int{
		{1, 2},
		{2, 5},
		{6, 9},
		{7, 11},
	}
	expect := []bool{true, false, true, true}
	runSample(t, nums, queries, expect)
}
