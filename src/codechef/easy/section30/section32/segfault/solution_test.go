package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, statements [][]int, expect []int) {
	res := solve(n, statements)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	statements := [][]int{
		{2, 2},
		{1, 1},
		{1, 3},
		{1, 3},
	}
	expect := []int{1, 2}
	runSample(t, n, statements, expect)
}
