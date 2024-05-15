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
	a := []int{2, 1}
	expect := []int{3, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 100, 1000, 1000000, 1000000000}
	expect := []int{1000000005, 1000004, 1003, 102, 2}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 4, 8}
	expect := []int{11, 7, 6}
	runSample(t, a, expect)
}
