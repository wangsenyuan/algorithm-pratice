package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, a int, b int, A [][]int, B [][]int, expect []int) {
	res := solve(k, a, b, A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, a, b := 10, 2, 1
	A := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	B := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	expect := []int{1, 9}
	runSample(t, k, a, b, A, B, expect)
}

func TestSample2(t *testing.T) {
	k, a, b := 8, 1, 1
	A := [][]int{
		{2, 2, 1},
		{3, 3, 1},
		{3, 1, 3},
	}
	B := [][]int{
		{1, 1, 1},
		{2, 1, 1},
		{1, 2, 3},
	}
	expect := []int{5, 2}
	runSample(t, k, a, b, A, B, expect)
}

func TestSample3(t *testing.T) {
	k, a, b := 5, 1, 1
	A := [][]int{
		{1, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	B := [][]int{
		{1, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	expect := []int{0, 0}
	runSample(t, k, a, b, A, B, expect)
}

func TestSample4(t *testing.T) {
	k, a, b := 10, 2, 1
	A := [][]int{
		{2, 2, 1},
		{3, 2, 2},
		{3, 1, 3},
	}
	B := [][]int{
		{3, 1, 3},
		{1, 2, 2},
		{3, 3, 2},
	}
	expect := []int{8, 1}
	runSample(t, k, a, b, A, B, expect)
}
