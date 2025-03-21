package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, b []int, queries [][]int, expect []int) {
	res := solve(a, b, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	b := []int{3, 1, 2, 1, 2}
	queires := [][]int{
		{3, 1, 3},
		{1, 2, 5, 3},
		{3, 1, 3},
		{1, 1, 3, 1},
		{2, 5, 5, 2},
		{3, 1, 5},
	}
	expect := []int{16, 25, 84}

	runSample(t, a, b, queires, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1000000000, 1000000000}
	b := []int{1000000000, 1000000000}
	queires := [][]int{
		{3, 1, 1},
		{1, 2, 2, 1000000000},
		{3, 1, 2},
	}
	expect := []int{716070898, 151723988}

	runSample(t, a, b, queires, expect)
}
