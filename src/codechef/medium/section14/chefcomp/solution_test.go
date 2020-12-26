package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, P []int, A []int, B []int, expect []int) {
	res := solve(n, E, P, A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	P := []int{1, 2, 3}
	A := []int{1, 2, 3}
	B := []int{1, 3, 6}
	expect := []int{1, 2, 3}
	runSample(t, n, E, P, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	P := []int{1, 2, 3}
	A := []int{5, 6, 3}
	B := []int{4, 10, 12}
	expect := []int{1, 2, -1}
	runSample(t, n, E, P, A, B, expect)
}
