package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x int, queries []int, expect []int) {
	res := solve(x, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 3
	queries := []int{0, 1, 2, 2, 0, 0, 10}
	expect := []int{1, 2, 3, 3, 4, 4, 7}
	runSample(t, x, queries, expect)
}

func TestSample2(t *testing.T) {
	x := 3
	queries := []int{1, 2, 1, 2}
	expect := []int{0, 0, 0, 0}
	runSample(t, x, queries, expect)
}
