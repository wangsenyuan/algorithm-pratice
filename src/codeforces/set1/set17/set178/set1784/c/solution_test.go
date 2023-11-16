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
	a := []int{3, 1, 2}
	expect := []int{2, 1, 0}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 1, 5, 4, 1, 1}
	expect := []int{3, 2, 4, 4, 4, 4}
	runSample(t, a, expect)
}
