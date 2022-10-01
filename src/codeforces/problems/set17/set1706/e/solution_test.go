package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, Q [][]int, expect []int) {
	res := solve(n, edges, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	Q := [][]int{
		{1, 1},
		{1, 2},
	}
	expect := []int{0, 1}
	runSample(t, n, edges, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 4},
		{3, 5},
	}
	Q := [][]int{
		{1, 4},
		{3, 4},
		{2, 2},
		{2, 5},
		{3, 5},
	}
	expect := []int{3, 3, 0, 5, 5}
	runSample(t, n, edges, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	Q := [][]int{
		{1, 3},
	}
	expect := []int{2}
	runSample(t, n, edges, Q, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 4},
		{1, 3},
		{2, 3},
	}
	Q := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 2},
		{2, 3},
		{2, 4},
		{3, 3},
		{3, 4},
		{4, 4},
	}
	expect := []int{0, 3, 3, 3, 0, 3, 3, 0, 2, 0}
	runSample(t, n, edges, Q, expect)
}
func TestSample5(t *testing.T) {
	n := 3
	edges := [][]int{
		{2, 1},
		{1, 3},
	}
	Q := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 2},
		{2, 3},
		{3, 3},
	}
	expect := []int{0, 1, 2, 0, 2, 0}
	runSample(t, n, edges, Q, expect)
}
