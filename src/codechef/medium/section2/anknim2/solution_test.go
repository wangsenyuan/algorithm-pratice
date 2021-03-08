package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int64) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 2, 3, 2, 1, 3}
	expect := []int64{0, 0, 3, 0, 0, 1}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 1, 1, 1}
	expect := []int64{0, 3, 0, 1}
	runSample(t, n, A, expect)
}
