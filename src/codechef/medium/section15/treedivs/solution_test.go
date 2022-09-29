package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, edges [][]int, expect []int) {
	res := solve(n, A, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{100, 101, 102, 103}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := []int{192, 2, 8, 2}
	runSample(t, n, A, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 2, 2, 2}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := []int{5, 4, 3, 2}
	runSample(t, n, A, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{43, 525, 524, 12, 289}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
	}
	expect := []int{1080, 12, 60, 18, 3}
	runSample(t, n, A, edges, expect)
}
