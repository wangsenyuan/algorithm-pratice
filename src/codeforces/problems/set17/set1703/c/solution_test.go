package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, moves []string, expect []int) {
	res := solve(n, A, moves)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{9, 3, 1}
	moves := []string{
		"DDD",
		"UDUU",
		"DU",
	}
	expect := []int{2, 1, 1}
	runSample(t, len(A), A, moves, expect)
}
