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
	a := []int{1, 3, 7, 5}
	q := []int{2, 10}
	expect := []int{2, 6}
	runSample(t, a, q, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 0, 2}
	q := []int{9, 8, 17, 1, 3}
	expect := []int{5, 4, 7, 0, 1}
	runSample(t, a, q, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10}
	q := []int{5, 2318381298321}
	expect := []int{15, 2318381298331}
	runSample(t, a, q, expect)
}
