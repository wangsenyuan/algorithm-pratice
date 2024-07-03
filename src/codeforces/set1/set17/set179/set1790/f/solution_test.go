package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, c0 int, c []int, edges [][]int, expect []int) {
	res := solve(n, c0, c, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, c0 := 6, 6
	c := []int{4, 1, 3, 5, 2}
	edges := [][]int{
		{2, 4},
		{6, 5},
		{5, 3},
		{3, 4},
		{1, 3},
	}
	expect := []int{3, 2, 1, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}

func TestSample2(t *testing.T) {
	n, c0 := 4, 2
	c := []int{4, 1, 3}
	edges := [][]int{
		{3, 1},
		{2, 3},
		{1, 4},
	}
	expect := []int{3, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}

func TestSample3(t *testing.T) {
	n, c0 := 10, 3
	c := []int{10, 7, 6, 5, 2, 9, 8, 1, 4}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{4, 5},
		{4, 3},
		{6, 4},
		{8, 7},
		{9, 8},
		{10, 8},
		{1, 8},
	}
	expect := []int{3, 2, 2, 2, 2, 2, 1, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}

func TestSample4(t *testing.T) {
	n, c0 := 7, 3
	c := []int{7, 5, 1, 2, 4, 6}
	edges := [][]int{
		{1, 2},
		{3, 2},
		{4, 5},
		{3, 4},
		{6, 5},
		{7, 6},
	}
	expect := []int{4, 2, 2, 1, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}

func TestSample5(t *testing.T) {
	n, c0 := 9, 7
	c := []int{9, 3, 1, 4, 2, 6, 8, 5}
	edges := [][]int{
		{4, 1},
		{8, 9},
		{4, 8},
		{2, 6},
		{7, 3},
		{2, 4},
		{3, 5},
		{5, 4},
	}
	expect := []int{5, 1, 1, 1, 1, 1, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}

func TestSample6(t *testing.T) {
	n, c0 := 10, 2
	c := []int{1, 8, 5, 10, 6, 9, 4, 3, 7}
	edges := [][]int{
		{10, 7},
		{7, 8},
		{3, 6},
		{9, 7},
		{7, 6},
		{4, 2},
		{1, 6},
		{7, 5},
		{9, 2},
	}
	expect := []int{4, 3, 2, 2, 1, 1, 1, 1, 1}
	runSample(t, n, c0, c, edges, expect)
}
