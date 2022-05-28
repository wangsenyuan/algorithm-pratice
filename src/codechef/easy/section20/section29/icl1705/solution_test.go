package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, mat [][]int, q int, F [][]int, S [][]int, expect []string) {
	res := solve(n, E, mat, q, F, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{5, 6},
		{5, 7},
		{7, 8},
	}
	mat := [][]int{
		{0, 1, 2},
		{1, 0, -2},
		{2, -2, 0},
	}
	q := 3
	F := [][]int{
		{3, 2},
		{2, 2},
		{1, 2},
	}
	S := [][]int{
		{3, 6, 4},
		{2, 3},
		{8},
	}
	expect := []string{
		"-2", "2", "No",
	}
	runSample(t, n, E, mat, q, F, S, expect)
}
