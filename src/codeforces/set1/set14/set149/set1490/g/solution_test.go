package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, X []int, expect []int64) {
	res := solve(A, X)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v,  but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, -3, 4}
	X := []int{1, 5, 2}
	expect := []int64{0, 6, 2}
	runSample(t, A, X, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-2, 0}
	X := []int{1, 2}
	expect := []int64{-1, -1}
	runSample(t, A, X, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 1}
	X := []int{1, 2}
	expect := []int64{1, 3}
	runSample(t, A, X, expect)
}
