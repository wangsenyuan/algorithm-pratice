package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, a []int, expect [][]int) {
	res := solve(n, a)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	a := []int{3, 5, 1, 4}
	expect := [][]int{
		{1, 3},
		{2, 5},
		{1, 4},
		{1, 5},
		{1, 5},
	}
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{1, 2, 4}
	expect := [][]int{
		{1, 3},
		{1, 2},
		{3, 4},
		{1, 4},
	}
	runSample(t, n, a, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	a := []int{10, 1, 5, 7, 1, 2, 5, 3, 6, 3, 9, 4, 3, 4, 9, 6, 8, 4, 9, 6}
	expect := [][]int{
		{1, 8},
		{1, 7},
		{1, 6},
		{1, 9},
		{1, 6},
		{1, 8},
		{1, 9},
		{1, 10},
		{1, 10},
		{1, 10},
	}
	runSample(t, n, a, expect)
}
