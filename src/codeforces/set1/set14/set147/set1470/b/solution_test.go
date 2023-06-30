package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q []int64, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 8, 4, 2}
	Q := []int64{0}
	expect := []int{2}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 5, 6, 2, 10, 7, 10, 4, 2, 10, 10, 6, 1, 1, 8, 1, 3, 7, 2, 9, 10, 4, 2, 4, 1, 9, 1, 1, 4, 1, 7, 8, 1, 7, 9, 3, 10, 7, 1}
	Q := []int64{1}
	expect := []int{31}
	runSample(t, A, Q, expect)
}
