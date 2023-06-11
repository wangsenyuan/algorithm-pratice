package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q []int, expect []bool) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	Q := []int{1, 8, 9, 12, 6}
	expect := []bool{true, false, true, false, true}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1, 3, 1, 3}
	Q := []int{1, 2, 3, 9, 11}
	expect := []bool{false, true, false, true, true}
	runSample(t, A, Q, expect)
}
