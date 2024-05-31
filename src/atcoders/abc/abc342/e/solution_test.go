package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, trains [][]int, expect []int) {
	res := solve(n, trains)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	trains := [][]int{
		{10, 5, 10, 3, 1, 3},
		{13, 5, 10, 2, 3, 4},
		{15, 5, 10, 7, 4, 6},
		{3, 10, 2, 4, 2, 5},
		{7, 10, 2, 3, 5, 6},
		{5, 3, 18, 2, 2, 3},
		{6, 3, 20, 4, 2, 1},
	}
	expect := []int{55, 56, 58, 60, 17}
	runSample(t, n, trains, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	trains := [][]int{
		{1000000000, 1000000000, 1000000000, 1000000000, 1, 5},
		{5, 9, 2, 6, 2, 3},
		{10, 4, 1, 6, 2, 3},
		{1, 1, 1, 1, 3, 5},
		{3, 1, 4, 1, 5, 1},
	}
	expect := []int{1000000000000000000, -1, 1, -1}
	runSample(t, n, trains, expect)
}
