package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, W []int, E [][]int, expect []int64) {
	res := solve(n, W, E)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	W := []int{3, 5, 4, 6}
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 3},
	}
	expect := []int64{18, 22, 25}
	runSample(t, n, W, E, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	W := []int{21, 32}
	E := [][]int{
		{2, 1},
	}
	expect := []int64{53}
	runSample(t, n, W, E, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	W := []int{20, 13, 17, 13, 13, 11}
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
	}
	expect := []int64{87, 107, 127, 147, 167}
	runSample(t, n, W, E, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	W := []int{10, 6, 6, 6}
	E := [][]int{
		{1, 2},
		{2, 3},
		{4, 1},
	}
	expect := []int64{28, 38, 44}
	runSample(t, n, W, E, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	W := []int{4, 5, 3, 6}
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 3},
	}
	expect := []int64{18, 22, 25}
	runSample(t, n, W, E, expect)
}
