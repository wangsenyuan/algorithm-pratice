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
	A := []int{3, 3, 3, 3, 3, 3}
	Q := []int64{18, 1}
	expect := []int{3, 3}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 2, 1}
	Q := []int64{1, 2, 3, 4}
	expect := []int{1, 1, 1, 2}
	runSample(t, A, Q, expect)
}
