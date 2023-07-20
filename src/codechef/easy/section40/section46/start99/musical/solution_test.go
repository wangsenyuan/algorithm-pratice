package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, C []int, L []int, A [][]int, expect []int) {
	res := solve(m, C, L, A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	C := []int{1, 2, 1}
	L := []int{5, 6, 1}
	A := [][]int{
		{3, 2},
		{1, 1},
		{2, 1},
	}
	expect := []int{500000004, 0, 0}
	runSample(t, m, C, L, A, expect)
}

func TestSample2(t *testing.T) {
	m := 8
	C := []int{4, 2, 8, 1, 1}
	L := []int{10, 2, 32, 8, 23, 7, 9}
	A := [][]int{
		{1, 3},
		{4, 5},
		{3, 2},
		{5, 1},
		{2, 8},
	}
	expect := []int{375000006, 750000009, 125000004, 500000006, 2}
	runSample(t, m, C, L, A, expect)
}
