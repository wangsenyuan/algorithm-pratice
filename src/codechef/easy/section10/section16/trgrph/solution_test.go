package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, edges [][]int, expect []int) {
	res := solve(n, m, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, m, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 0
	var edges [][]int
	expect := []int{0}
	runSample(t, n, m, edges, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 2
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := []int{2, 0, 2}
	runSample(t, n, m, edges, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	expect := []int{0, 1, 2}
	runSample(t, n, m, edges, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 2
	edges := [][]int{
		{1, 4},
		{2, 3},
	}
	var expect []int
	runSample(t, n, m, edges, expect)
}

func TestSample5(t *testing.T) {
	n, m := 7, 10
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
	}
	expect := []int{0, 1, 1, 2, 2, 3, 3}
	runSample(t, n, m, edges, expect)
}
