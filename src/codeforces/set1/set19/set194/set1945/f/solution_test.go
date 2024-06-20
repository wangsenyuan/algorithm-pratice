package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, p []int, expect []int) {
	res := solve(a, p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{9, 8, 14}
	p := []int{3, 2, 1}
	expect := []int{16, 2}
	runSample(t, a, p, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	p := []int{1, 2, 3, 4, 5}
	expect := []int{9, 3}
	runSample(t, a, p, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	p := []int{6, 5, 4, 3, 2, 1}
	expect := []int{8, 2}
	runSample(t, a, p, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 2, 5, 5}
	p := []int{4, 2, 3, 1}
	expect := []int{5, 1}
	runSample(t, a, p, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 2, 9, 10, 10}
	p := []int{1, 4, 2, 3, 5}
	expect := []int{20, 2}
	runSample(t, a, p, expect)
}
