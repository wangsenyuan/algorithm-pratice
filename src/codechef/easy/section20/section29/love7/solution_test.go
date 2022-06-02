package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q []int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 4, 5, 6}
	Q := []int{3, 5}
	expect := []int{3, 0}
	runSample(t, A, Q, expect)
}
