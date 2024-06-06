package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, a []int, b []int, d []int, expect []int) {
	res := solve(n, a, b, d)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	a := []int{10}
	b := []int{10}
	d := []int{0, 6, 7}
	expect := []int{0, 6, 7}
	runSample(t, n, a, b, d, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	a := []int{4, 10}
	b := []int{4, 7}
	d := []int{6, 4, 2, 7}
	expect := []int{5, 4, 2, 5}
	runSample(t, n, a, b, d, expect)
}
