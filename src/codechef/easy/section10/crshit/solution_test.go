package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, K int64, D []int, X []int64, T []int64, expect []int64) {
	res := solve(N, Q, K, D, X, T)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 5, 3
	var K int64 = 11
	D := []int{1, 1, 2, 2, 2}
	X := []int64{3, 10, 4, 7, 0}
	T := []int64{3, 8, 100}
	expect := []int64{4, 10, 110}
	runSample(t, N, Q, K, D, X, T, expect)
}
