package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, C []int, expect []int) {
	res := solve(n, P, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	P := []int{0, 1, 2, 3, 3, 1, 6, 6, 6, 1}
	C := []int{5, 6, 3, 4, 2, 1, 7, 8, 9, 10}
	expect := []int{4, 1, 2, 2, 0, 0, 1, 1, 1, 1}
	runSample(t, n, P, C, expect)
}
