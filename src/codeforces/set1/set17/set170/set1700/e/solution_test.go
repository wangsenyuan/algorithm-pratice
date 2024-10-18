package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a [][]int, expect []int) {
	res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{2, 1, 3},
		{6, 7, 4},
		{9, 8, 5},
	}
	expect := []int{0}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]int{
		{1, 6, 4},
		{3, 2, 5},
	}
	expect := []int{1, 3}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := [][]int{
		{1, 6, 5, 4, 3, 2},
	}
	expect := []int{2}
	runSample(t, a, expect)
}
