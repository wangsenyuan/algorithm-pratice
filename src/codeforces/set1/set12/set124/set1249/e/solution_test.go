package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, c int, a []int, b []int, expect []int) {
	res := solve(n, c, a, b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, c := 10, 2
	a := []int{7, 6, 18, 6, 16, 18, 1, 17, 17}
	b := []int{6, 9, 3, 10, 9, 1, 10, 1, 5}
	expect := []int{0, 7, 13, 18, 24, 35, 36, 37, 40, 45}

	runSample(t, n, c, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, c := 10, 1
	a := []int{3, 2, 3, 1, 3, 3, 1, 4, 1}
	b := []int{1, 2, 3, 4, 4, 1, 2, 1, 3}
	expect := []int{0, 2, 4, 7, 8, 11, 13, 14, 16, 17}

	runSample(t, n, c, a, b, expect)
}
