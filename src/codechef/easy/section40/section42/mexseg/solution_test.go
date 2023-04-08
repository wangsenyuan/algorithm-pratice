package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, Q [][]int, expect []int64) {
	res := solve(P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{2, 0, 1, 3}
	Q := [][]int{
		{2, 3, 0, 2},
		{4, 4, 3, 4},
	}
	expect := []int64{4, 1}
	runSample(t, P, Q, expect)
}
