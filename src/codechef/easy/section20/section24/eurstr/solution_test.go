package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S []string, Q [][]int, expect []int) {
	res := solve(n, S, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	S := []string{
		"abac",
		"abab",
	}
	Q := [][]int{
		{2, 2},
		{1, 1},
	}
	expect := []int{4, 6}
	runSample(t, n, S, Q, expect)
}
