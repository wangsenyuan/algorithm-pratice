package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, b []int, expect []int) {
	res := solve(a, b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 7, 3, 4, 4, 6, 5}
	b := []int{6, 3, 4, 4, 7, 6, 5}
	expect := []int{2, 5}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1}
	b := []int{1, 1, 2}
	expect := []int{1, 3}
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 2, 1}
	b := []int{2, 1, 2}
	expect := []int{2, 3}
	runSample(t, a, b, expect)
}
