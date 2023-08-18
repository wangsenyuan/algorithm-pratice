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
	A := []int{6, 2, 3, 7}
	Q := [][]int{
		{3, 2},
		{4, 7},
	}
	expect := []int{3, 4}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 13, 5, 2, 6}
	Q := [][]int{
		{2, 4},
		{5, 2},
		{1, 5},
		{3, 2},
	}
	expect := []int{4, 4, 2, 3}
	runSample(t, A, Q, expect)
}

func TestSample3(t *testing.T) {
	A := []int{769, 514, 336, 173, 181, 373, 519, 338, 985, 709, 729, 702, 168}
	Q := [][]int{
		{12, 581},
		{6, 222},
		{7, 233},
		{5, 117},
	}
	expect := []int{5, 6, 6, 5}
	runSample(t, A, Q, expect)
}
