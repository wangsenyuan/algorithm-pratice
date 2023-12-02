package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q [][]int, expect []int) {
	res := solve(Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := [][]int{
		{1, 3, 2, 5},
		{2, 4, 1},
		{2, 3, 2},
	}
	expect := []int{1, 2, 5}
	runSample(t, Q, expect)
}

func TestSample2(t *testing.T) {
	Q := [][]int{
		{1, 6, 5, 1},
		{2, 3, 1},
		{2, 6, 2},
	}
	expect := []int{1, -1, 1}
	runSample(t, Q, expect)
}

func TestSample3(t *testing.T) {
	Q := [][]int{
		{1, 4, 2, 2},
		{1, 2, 1, 3},
		{2, 10, 2},
	}
	expect := []int{1, 0, 1}
	runSample(t, Q, expect)
}

func TestSample4(t *testing.T) {
	Q := [][]int{
		{1, 7, 3, 6},
		{1, 2, 1, 8},
		{2, 5, 1},
		{1, 10, 9, 7},
		{1, 8, 1, 2},
		{1, 10, 5, 8},
		{1, 10, 7, 7},
		{2, 7, 4},
		{1, 9, 4, 2},
	}
	expect := []int{1, 0, -1, 0, 0, 0, 1, 8, 0}
	runSample(t, Q, expect)
}
func TestSample5(t *testing.T) {
	Q := [][]int{
		{1, 2, 1, 6},
		{1, 8, 5, 6},
		{1, 4, 2, 7},
		{2, 9, 1},
		{1, 5, 1, 4},
		{1, 5, 2, 7},
		{1, 7, 1, 9},
		{1, 9, 1, 4},
		{2, 10, 8},
	}
	expect := []int{1, 0, 0, 1, 0, 0, 0, 0, 1}
	runSample(t, Q, expect)
}
