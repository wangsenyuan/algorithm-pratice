package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect [][]int) {
	res := solve(n, E, Q)

	for i := 0; i < len(Q); i++ {
		if !reflect.DeepEqual(res[i], expect[i]) {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	Q := [][]int{
		{1, 3},
		{1, 2},
	}
	expect := [][]int{
		{2},
		{1, 2},
	}
	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
	}
	Q := [][]int{
		{2},
		{4, 5},
	}
	expect := [][]int{
		{2},
		{2, 3},
	}
	runSample(t, n, E, Q, expect)
}
