package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q []int, expect []int64) {
	res := solve(n, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{-3, 5, -3, 2, 8, -20, 6, -1}
	Q := []int{1, 0, 1}
	expect := []int64{18, 8, 13}
	runSample(t, n, A, Q, expect)
}
