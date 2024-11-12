package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, x []int, expect []int) {
	res := solve(n, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	x := []int{1, 2, 3, 4}
	expect := []int{3, 4, 6, 5}
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	x := []int{2, 1, 5, 3, 5}
	expect := []int{9, 8, 12, 6, 8}
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	x := []int{1, 2, 1, 1, 2, 2, 2, 2, 2, 2}
	expect := []int{3, 3}
	runSample(t, n, x, expect)
}
