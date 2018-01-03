package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, nums []int, query [][]int, expect []int) {
	res := solve(n, nums, query)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample: %d, %v, %v, should give answer %v, but got %v", n, nums, query, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	nums := []int{2, 1, 4, 3}
	query := [][]int{
		{1, 1, 4},
		{2, 2, 3},
		{1, 1, 3},
	}
	expect := []int{0, 1}
	runSample(t, n, nums, query, expect)
}
