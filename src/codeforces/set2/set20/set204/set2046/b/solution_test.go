package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 3}
	expect := []int{1, 3, 3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2, 1, 4}
	expect := []int{1, 1, 3, 3, 5}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 6, 5, 4}
	expect := []int{1, 2, 3, 4, 6, 7}
	runSample(t, a, expect)
}
