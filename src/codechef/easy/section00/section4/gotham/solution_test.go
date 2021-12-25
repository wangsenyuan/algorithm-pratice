package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, Q [][]int, expect []int64) {
	res := solve(len(P), P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{5, 5, 6, 1}
	Q := [][]int{
		{2, 11},
		{2, 3},
	}
	expect := []int64{6, 2}
	runSample(t, P, Q, expect)
}
