package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A [][]int, expect [][]int) {
	solve(A)

	if !reflect.DeepEqual(A, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, A)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{2, 1, 2},
		{2, 1, 2},
		{1, 1, 2},
	}
	expect := [][]int{
		{2, 1, 1},
		{2, 1, 1},
		{2, 2, 2},
	}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{3, 3, 1, 2},
		{1, 1, 3, 1},
		{3, 2, 3, 2},
		{2, 3, 3, 1},
	}
	expect := [][]int{
		{3, 1, 1, 2},
		{3, 1, 2, 1},
		{3, 3, 3, 3},
		{2, 3, 2, 1},
	}
	runSample(t, A, expect)
}
