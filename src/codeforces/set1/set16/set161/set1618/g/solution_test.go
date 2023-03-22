package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, K []int, expect []int64) {
	res := solve(A, B, K)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 30, 15}
	B := []int{12, 31, 14, 18}
	K := []int{0, 1, 2, 3, 4}
	expect := []int64{55, 56, 60, 64, 64}
	runSample(t, A, B, K, expect)
}
