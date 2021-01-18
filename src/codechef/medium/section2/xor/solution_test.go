package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, k int, expect []int) {
	res := solve(n, k, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 5
	A := []int{1, 1, 3, 4}
	expect := []int{0, 2, 2, 5, 5}
	runSample(t, n, A, k, expect)
}
