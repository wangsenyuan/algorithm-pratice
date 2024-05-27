package main

import (
	"reflect"
	"strconv"
	"testing"
)

func runSample(t *testing.T, n int, X []int, V []int, queries [][]int, expect []int) {
	res := solve(n, X, V, queries)

	// ss := toStrings(expect)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func toStrings(nums []int) []string {
	res := make([]string, len(nums))

	for i, num := range nums {
		res[i] = strconv.Itoa(num)
	}

	return res
}

func TestSample1(t *testing.T) {
	n := 8
	x := []int{1, 3, 8}
	v := []int{3, 24, 10}
	queies := [][]int{
		{2, 2, 5},
		{1, 5, 15},
		{2, 5, 5},
		{2, 7, 8},
	}
	expect := []int{171, 0, 15}
	runSample(t, n, x, v, queies, expect)
}
