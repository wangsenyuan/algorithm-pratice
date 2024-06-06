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
	a := []int{1, 0, 3, 2}
	expect := []int{0, 1, 2, 3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 7, 1, 5, 6}
	expect := []int{1, 5, 2, 6, 7}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{16, 4, 1, 64}
	expect := []int{16, 4, 1, 64}
	runSample(t, a, expect)
}
