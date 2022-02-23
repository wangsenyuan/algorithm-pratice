package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, edges int, Q []int, expect []int64) {
	res := solve(n, A, edges, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{8, 9, 6, 10}
	edges := 3
	Q := []int{
		2, 0, 0, 1, 0, 2,
	}
	expect := []int64{33, 37}
	runSample(t, n, A, edges, Q, expect)
}
