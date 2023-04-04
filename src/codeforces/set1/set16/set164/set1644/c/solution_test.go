package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, x int, expect []int64) {
	res := solve(A, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 1, 3, 2}
	x := 2
	expect := []int64{10, 12, 14, 16, 18}
	runSample(t, A, x, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-2, -7, -1}
	x := 5
	expect := []int64{0, 4, 4, 5}
	runSample(t, A, x, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-6, -1, -2, 4, -6, -1, -4, 4, -5, -4}
	x := 2
	expect := []int64{4, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8}
	runSample(t, A, x, expect)
}
