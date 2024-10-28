package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect [][]int) {
	res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1, 2, 1}
	expect := [][]int{
		{1, 3},
		{3, 1},
	}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := [][]int{
		{1, 4},
		{2, 2},
		{4, 1},
	}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 1, 2}

	runSample(t, a, nil)
}

func TestSample4(t *testing.T) {
	a := []int{2, 1, 2, 1, 1, 1, 1, 1}
	expect := [][]int{
		{1, 6},
		{2, 3},
		{6, 1},
	}
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 1, 2, 2, 2, 1, 1, 2, 2, 1, 1, 1, 2, 1, 1, 2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 1, 2, 1, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1, 1, 1, 2, 2, 1, 1, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 2, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 2, 1, 1, 2, 2, 1, 2, 2, 1, 1, 1, 1, 1, 2, 2, 1, 1, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 1, 2, 2, 1, 1, 1, 2, 2, 2, 1, 2, 2, 1, 1, 2, 1, 1, 1, 1, 2, 1, 1}
	expect := [][]int{
		{1, 100},
		{2, 50},
		{6, 11},
		{8, 8},
		{19, 4},
		{25, 3},
		{40, 2},
		{100, 1},
	}
	runSample(t, a, expect)
}
