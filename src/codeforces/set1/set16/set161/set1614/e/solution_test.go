package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, T []int, X [][]int, expect []int) {
	res := solve(n, T, X)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	T := []int{50, 50, 0}
	X := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{2, 5, 9, 15, 22, 30, 38, 47, 53}
	runSample(t, n, T, X, expect)
}
