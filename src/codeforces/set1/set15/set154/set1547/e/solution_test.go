package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, a []int, x []int, expect []int) {
	res := solve(n, a, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	a := []int{2, 5}
	x := []int{14, 16}
	expect := []int{15, 14, 15, 16, 16, 17}
	runSample(t, n, a, x, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	a := []int{7}
	x := []int{30}
	expect := []int{36, 35, 34, 33, 32, 31, 30, 31, 32, 33}
	runSample(t, n, a, x, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	a := []int{3, 1, 4, 2, 5}
	x := []int{3, 1, 4, 2, 5}
	expect := []int{1, 2, 3, 4, 5}
	runSample(t, n, a, x, expect)
}
