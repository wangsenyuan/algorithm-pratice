package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, X []int, expect []int) {
	res := solve(A, X)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 9, 3, 15, 10}
	X := []int{2, 2, 3, 7}
	expect := []int{2, 10, 3, 9}
	runSample(t, A, X, expect)
}
