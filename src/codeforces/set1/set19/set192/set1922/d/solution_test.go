package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, d []int, expect []int) {
	res := solve(a, d)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 4, 7, 5, 10}
	d := []int{4, 9, 1, 18, 1}
	expect := []int{3, 1, 0, 0, 0}
	runSample(t, a, d, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1}
	d := []int{1, 3}
	expect := []int{0, 0}
	runSample(t, a, d, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 2, 4}
	d := []int{3, 3, 4, 2}
	expect := []int{1, 1, 1, 0}
	runSample(t, a, d, expect)
}
