package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, routes [][]int, expect []int) {
	res := solve(a, routes)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	routes := [][]int{
		{1, 2, 4},
		{2, 3, 7},
	}
	a := []int{6, 20, 1, 25}

	expect := []int{6, 14, 1, 25}

	runSample(t, a, routes, expect)
}

func TestSample2(t *testing.T) {
	routes := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{1, 3, 1},
	}
	a := []int{30, 10, 20}

	expect := []int{12, 10, 12}

	runSample(t, a, routes, expect)
}
