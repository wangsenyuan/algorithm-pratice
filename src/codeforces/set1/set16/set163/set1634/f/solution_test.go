package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, mod int, Q []string, expect []bool) {
	res := solve(A, B, mod, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2, 1}
	B := []int{0, 0, 0}
	mod := 3
	Q := []string{
		"A 1 3",
		"A 1 3",
		"B 1 1",
		"B 2 2",
		"A 3 3",
	}
	expect := []bool{true, false, false, false, true}

	runSample(t, A, B, mod, Q, expect)
}
