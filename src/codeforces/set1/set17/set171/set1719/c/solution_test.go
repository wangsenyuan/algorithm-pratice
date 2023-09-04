package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 2}
	Q := [][]int{{1, 2}}
	expect := []int{2}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 4, 2}
	Q := [][]int{
		{4, 5},
		{3, 2},
	}
	expect := []int{0, 1}
	runSample(t, A, Q, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 5, 4}
	Q := [][]int{
		{5, 1000000000},
		{4, 6},
	}
	expect := []int{0, 4}
	runSample(t, A, Q, expect)
}
