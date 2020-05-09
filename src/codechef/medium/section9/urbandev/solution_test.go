package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, lines [][]int, expect []int) {
	_, res := solve(n, lines)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	lines := [][]int{
		{2, 1, 2, 6},
		{1, 3, 5, 3},
		{4, 1, 4, 4},
		{1, 5, 5, 5},
	}
	expect := []int{2, 2, 1, 1}
	runSample(t, n, lines, expect)
}
