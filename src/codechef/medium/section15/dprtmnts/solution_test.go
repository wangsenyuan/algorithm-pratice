package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, conns [][]int, expect_deps []int, expect_heads []int) {
	_, deps, heads := solve(n, m, conns)

	if !reflect.DeepEqual(deps, expect_deps) || !reflect.DeepEqual(heads, expect_heads) {
		t.Errorf("Sample expect %v\n, %v\n, but got %v\n, %v\n", expect_deps, expect_heads, deps, heads)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	m := 8
	conns := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{4, 5},
		{6, 7},
		{2, 4},
		{4, 6},
		{2, 6},
	}
	deps := []int{1, 1, 1, 2, 2, 3, 3}
	heads := []int{2, 4, 6}
	runSample(t, n, m, conns, deps, heads)
}

func TestSample2(t *testing.T) {
	n := 2
	m := 1
	conns := [][]int{
		{1, 2},
	}
	deps := []int{1, 1}
	heads := []int{1}
	runSample(t, n, m, conns, deps, heads)
}

func TestSample3(t *testing.T) {
	n := 6
	m := 8
	conns := [][]int{
		{6, 2},
		{5, 4},
		{3, 1},
		{5, 2},
		{3, 6},
		{2, 1},
		{6, 1},
		{2, 3},
	}
	deps := []int{1, 1, 1, 2, 2, 1}
	heads := []int{2, 5}
	runSample(t, n, m, conns, deps, heads)
}
