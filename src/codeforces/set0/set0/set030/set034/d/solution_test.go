package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, r1 int, r2 int, p []int, expect []int) {
	res := solve(n, r1, r2, p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, r1, r2 := 3, 2, 3
	p := []int{2, 2}
	expect := []int{2, 3}
	runSample(t, n, r1, r2, p, expect)
}

func TestSample2(t *testing.T) {
	n, r1, r2 := 6, 2, 4
	p := []int{6, 1, 2, 4, 2}
	expect := []int{6, 4, 1, 4, 2}
	runSample(t, n, r1, r2, p, expect)
}
