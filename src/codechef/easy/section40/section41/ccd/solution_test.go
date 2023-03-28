package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B string, Q [][]int, expect []bool) {
	res := solve(A, B, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "abcdz"
	B := "nbhea"
	Q := [][]int{
		{2, 2},
		{1, 3},
		{4, 5},
	}
	expect := []bool{true, false, true}

	runSample(t, A, B, Q, expect)
}
