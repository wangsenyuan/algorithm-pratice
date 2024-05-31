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
	a := []int{4}
	expect := []int{3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 2, 5}
	expect := []int{3, 2, 2, 3}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 5, 2, 3}
	expect := []int{3, 1, 3, 3}
	runSample(t, a, expect)
}
