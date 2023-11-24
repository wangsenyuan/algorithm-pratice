package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, k int, expect []int) {
	res := solve(a, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2, 3, 3}
	k := 3
	expect := []int{1, 3, 2}
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3, 3, 4, 4, 2}
	k := 4
	expect := []int{4, -(inf + 1), 3}
	runSample(t, a, k, expect)
}
