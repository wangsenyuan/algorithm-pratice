package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, p []int, expect []int) {
	res := solve(n, p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	p := []int{1, 3, 4, 2}
	expect := []int{1, 2, 3, 2, 1}
	runSample(t, n, p, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	p := []int{6, 8, 3, 4, 7, 2, 1, 5}
	expect := []int{1, 2, 2, 3, 4, 3, 4, 5, 1}
	runSample(t, n, p, expect)
}
