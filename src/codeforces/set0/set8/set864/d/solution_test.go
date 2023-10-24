package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	_, res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 2, 3}
	expect := []int{1, 2, 4, 3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 5, 6, 3, 2, 1}
	expect := []int{4, 5, 6, 3, 2, 1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 8, 4, 6, 7, 1, 6, 3, 4, 5}
	expect := []int{2, 8, 4, 6, 7, 1, 9, 3, 10, 5}
	runSample(t, a, expect)
}
