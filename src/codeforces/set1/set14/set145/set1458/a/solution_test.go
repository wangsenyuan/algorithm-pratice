package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int64, B []int64, expect []int64) {
	res := solve(A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 25, 121, 169}
	B := []int64{1, 2, 7, 23}
	expect := []int64{2, 3, 8, 24}
	runSample(t, A, B, expect)
}
