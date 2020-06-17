package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{3, 2, 5, 4, 1, 7}
	expect := []int{0, 1, 0, 2, 1, 0}
	runSample(t, n, A, expect)
}
