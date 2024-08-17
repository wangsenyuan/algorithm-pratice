package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, q []int, expect [][]int) {
	res := solve(q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	q := []int{3, 3, 4, 4, 7, 7, 7}
	expect := [][]int{
		{3, 1, 4, 2, 7, 5, 6},
		{3, 2, 4, 1, 7, 6, 5},
	}
	runSample(t, q, expect)
}

func TestSample2(t *testing.T) {
	q := []int{1, 2, 3, 4}
	expect := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	runSample(t, q, expect)
}

func TestSample3(t *testing.T) {
	q := []int{3, 4, 5, 5, 5, 7, 7}
	expect := [][]int{
		{3, 4, 5, 1, 2, 7, 6},
		{3, 4, 5, 2, 1, 7, 6},
	}
	runSample(t, q, expect)
}
