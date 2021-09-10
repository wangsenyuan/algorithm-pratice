package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, B int, N []int, Q []int, expect []int) {
	res := solve(B, N, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := 3
	N := []int{2, 1, 0}
	Q := []int{1, 2}
	expect := []int{2, 1}
	runSample(t, B, N, Q, expect)
}
