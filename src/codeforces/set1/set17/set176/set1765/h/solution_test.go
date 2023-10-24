package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int, a []int, b []int, expect []int) {
	res := solve(p, a, b)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 3, 2, 4}
	a := []int{3}
	b := []int{1}
	expect := []int{2, 3, 1, 4}
	runSample(t, p, a, b, expect)
}

func TestSample2(t *testing.T) {
	p := []int{3, 3, 3}
	//a := []int{3}
	//b := []int{1}
	expect := []int{1, 1, 1}
	runSample(t, p, nil, nil, expect)
}

func TestSample3(t *testing.T) {
	p := []int{4, 3, 3, 2, 5}
	a := []int{3, 1, 4}
	b := []int{1, 5, 2}
	expect := []int{4, 2, 1, 1, 5}
	runSample(t, p, a, b, expect)
}
