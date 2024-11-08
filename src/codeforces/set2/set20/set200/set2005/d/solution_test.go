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
	a := []int{11, 4, 16, 17, 3, 24, 25, 8}
	b := []int{8, 10, 4, 21, 17, 18, 25, 21}
	expect := []int{2, 36}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 4, 24, 13}
	b := []int{15, 3, 1, 14}
	expect := []int{3, 2}
	runSample(t, a, b, expect)
}
