package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, c []int, edges [][]int, queries []int, expect []bool) {
	res := solve(n, c, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	c := []int{1, 0}
	edges := [][]int{
		{1, 2},
	}
	queries := []int{1}
	expect := []bool{false}
	runSample(t, n, c, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	c := []int{1, 0, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 5},
		{3, 4},
	}
	queries := []int{4, 3, 2, 5}
	expect := []bool{false, true, true, false}
	runSample(t, n, c, edges, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	c := []int{1, 1, 1, 1, 1}
	edges := [][]int{
		{3, 5},
		{2, 5},
		{3, 4},
		{1, 5},
	}
	queries := []int{1, 1, 1}
	expect := []bool{true, false, true}
	runSample(t, n, c, edges, queries, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	c := []int{0, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
	}
	queries := []int{1, 2, 3, 2}
	expect := []bool{true, true, true, false}
	runSample(t, n, c, edges, queries, expect)
}

func TestSample5(t *testing.T) {
	n := 1
	c := []int{0}

	queries := []int{1}
	expect := []bool{true}
	runSample(t, n, c, nil, queries, expect)
}

func TestSample6(t *testing.T) {
	n := 5
	c := []int{0, 0, 0, 0, 0}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}

	queries := []int{1, 2, 3, 4, 5, 2, 3, 4, 5}
	expect := []bool{true, true, true, true, true, false, false, false, true}
	runSample(t, n, c, edges, queries, expect)
}
