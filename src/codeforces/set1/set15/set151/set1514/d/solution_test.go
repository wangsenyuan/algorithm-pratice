package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{1, 3, 2, 3, 3, 2}
	queries := [][]int{
		{1, 6},
		{2, 5},
	}
	expect := []int{1, 2}
	runSample(t, a, queries, expect)
}
