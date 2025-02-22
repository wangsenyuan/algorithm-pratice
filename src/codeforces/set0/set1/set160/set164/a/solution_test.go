package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, V []int, E [][]int, expect []int) {
	res := solve(n, V, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 4
	V := []int{1, 0, 0, 2}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := []int{1, 1, 1, 1}
	runSample(t, n, V, E, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	V := []int{1, 0, 2}
	E := [][]int{
		{1, 3},
	}
	expect := []int{1, 0, 1}
	runSample(t, n, V, E, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	V := []int{2, 0, 1}
	E := [][]int{
		{1, 3},
	}
	expect := []int{0, 0, 0}
	runSample(t, n, V, E, expect)
}
