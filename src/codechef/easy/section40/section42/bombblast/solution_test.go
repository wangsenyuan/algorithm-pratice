package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, X []int, Y []int, expect []int64) {
	res := solve(n, m, X, Y)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 2
	X := []int{3, 1}
	Y := []int{2, 4}
	expect := []int64{-6, 0, 6, 14, 26}
	runSample(t, n, m, X, Y, expect)
}

func TestSample2(t *testing.T) {
	n, m := 10, 3
	X := []int{4, 7, 5}
	Y := []int{4, 5, 3}
	expect := []int64{-12, -48, -28, -14, -4, 6, 18, 34, 6, 12}
	runSample(t, n, m, X, Y, expect)
}
