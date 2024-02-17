package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, x []int, expect []int) {
	res := solve(a, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 4}
	x := []int{2, 3, 4}
	expect := []int{1, 2, 3, 6, 6}
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 8, 12, 36, 48, 6, 3}
	x := []int{10, 4, 2}
	expect := []int{7, 10, 14, 38, 58, 6, 3}
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 4, 8, 16}
	x := []int{5, 2, 3, 4, 1}
	expect := []int{1, 3, 7, 11, 19}
	runSample(t, a, x, expect)
}
