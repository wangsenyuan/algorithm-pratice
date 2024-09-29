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
	a := []int{1, 2, 3, 1, 2, 3}
	expect := []int{1, 2, 3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 5, 6, 5, 6, 7, 7}
	expect := []int{7}
	runSample(t, a, expect)
}
