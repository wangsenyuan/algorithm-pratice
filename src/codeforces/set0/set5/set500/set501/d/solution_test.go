package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int, q []int, expect []int) {
	res := solve(p, q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{0, 1}
	q := []int{0, 1}
	expect := []int{0, 1}
	runSample(t, p, q, expect)
}

func TestSample2(t *testing.T) {
	p := []int{0, 1}
	q := []int{1, 0}
	expect := []int{1, 0}
	runSample(t, p, q, expect)
}

func TestSample3(t *testing.T) {
	p := []int{1, 2, 0}
	q := []int{2, 1, 0}
	expect := []int{1, 0, 2}
	runSample(t, p, q, expect)
}
