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
	a := []int{11, 5, 2, 7}
	expect := []int{1, 1, 3, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 5, 3}
	expect := []int{5, 4, 3, 6}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 3, 6}
	expect := []int{-1, -1, -1}
	runSample(t, a, expect)
}
