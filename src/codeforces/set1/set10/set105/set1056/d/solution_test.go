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
	n := 3
	p := []int{1, 1}
	expect := []int{1, 1, 2}
	runSample(t, n, p, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	p := []int{1, 1, 3, 3}
	expect := []int{1, 1, 1, 2, 3}
	runSample(t, n, p, expect)
}
