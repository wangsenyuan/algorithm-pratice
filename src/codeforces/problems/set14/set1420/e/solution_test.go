package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 0, 0, 0, 1}
	expect := []int{0, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 12
	A := []int{0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0}
	expect := []int{9, 12, 13, 14, 14, 14, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15}
	runSample(t, n, A, expect)
}
