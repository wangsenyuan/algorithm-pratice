package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, B []int, expect []int) {
	res := solve(n, A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{4, 1, 9}
	B := []int{3, 4, 1}
	expect := []int{280, 204, 72}
	runSample(t, n, A, B, expect)
}
