package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 3},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}
	expect := []int{1, 4}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 3},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 6},
	}
	expect := []int{3, 3}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{2, 7},
		{3, 4},
		{4, 7},
		{5, 6},
		{6, 7},
	}
	expect := []int{1, 6}
	runSample(t, n, edges, expect)
}
