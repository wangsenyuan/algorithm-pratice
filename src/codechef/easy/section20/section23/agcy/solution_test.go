package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q [][]int, expect []int64) {
	res := solve(n, Q)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	Q := [][]int{
		{1, 3},
		{1, 2},
		{4, 5},
	}
	expect := []int64{2, 4, 3, 1, 2}
	runSample(t, n, Q, expect)
}
