package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []uint64, B []uint64, Q []int, expect []uint64) {
	res := solve(n, A, B, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []uint64{1, 2, 3}
	B := []uint64{4, 5, 6}
	Q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []uint64{5, 6, 6, 7, 7, 7, 8, 8, 9}
	runSample(t, n, A, B, Q, expect)
}
