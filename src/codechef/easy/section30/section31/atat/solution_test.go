package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect [][]int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{1, 0}
	expect := [][]int{
		{1, 2},
		{0, 2},
	}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{3, 1, 5, 0, 2}
	expect := [][]int{
		{4, 13},
		{3, 13},
		{2, 13},
		{1, 13},
		{1, 15},
	}
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{1, 2, 1}
	expect := [][]int{
		{3, 8},
		{1, 2},
		{0, 4},
	}
	runSample(t, n, A, expect)
}
