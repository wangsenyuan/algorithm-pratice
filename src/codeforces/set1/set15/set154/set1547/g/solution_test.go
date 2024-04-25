package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, routes [][]int, expect []int) {
	res := solve(n, routes)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample %d %v, expect %v, but got %v", n, routes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	routes := [][]int{
		{1, 4},
		{1, 3},
		{3, 4},
		{4, 5},
		{2, 1},
		{5, 5},
		{5, 6},
	}
	expect := []int{1, 0, 1, 2, -1, -1}
	runSample(t, n, routes, expect)
}

func TestSample2(t *testing.T) {
	n := 1

	expect := []int{1}
	runSample(t, n, nil, expect)
}

func TestSample3(t *testing.T) {
	n := 3

	routes := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}

	expect := []int{-1, -1, -1}
	runSample(t, n, routes, expect)
}

func TestSample4(t *testing.T) {
	n := 4

	routes := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 3},
	}

	expect := []int{1, 1, 2, 1}
	runSample(t, n, routes, expect)
}

func TestSample5(t *testing.T) {
	n := 5

	expect := []int{1, 0, 0, 0, 0}
	runSample(t, n, nil, expect)
}
