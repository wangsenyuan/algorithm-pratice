package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, A []int, expect []int) {
	res := solve(n, k, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 8, 3
	A := []int{1, 2, 2, 1, 1, 1, 2, 3}
	expect := []int{1, 1, 3}
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 2
	A := []int{1, 2, 2, 1, 1, 1, 2, 2}
	expect := []int{0, 0}
	runSample(t, n, k, A, expect)
}
