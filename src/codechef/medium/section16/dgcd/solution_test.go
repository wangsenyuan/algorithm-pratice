package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, E [][]int, commands [][]int, expect []int) {
	res := solve(n, E, A, commands)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{0, 4},
		{0, 5},
		{1, 5},
		{5, 2},
		{3, 5},
	}
	A := []int{3, 1, 3, 2, 4, 2}
	commands := [][]int{
		{3, 5},
		{1, 3, 1},
		{3, 4, 4},
		{3, 0},
		{2, 5},
	}
	expect := []int{2, 7, 1}
	runSample(t, n, A, E, commands, expect)
}
