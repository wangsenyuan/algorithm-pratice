package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, q []int, expect []int) {
	res := solve(a, q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 2, 8, 4}
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{3, 4, 5, 6, 7, 8, 8, 10, 8, 12}
	runSample(t, a, q, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 2, 8, 4, 4}
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{3, 4, 5, 6, 7, 8, 9, 8, 11, 8}
	runSample(t, a, q, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3}
	q := []int{10, 6, 8, 1, 3}
	expect := []int{10, 7, 8, 3, 3}
	runSample(t, a, q, expect)
}
