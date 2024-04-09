package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, a []int, queries []string, expect []int) {
	res := solve(n, P, a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	P := []int{1, 1, 1}
	a := []int{1, 0, 0, 1}
	queries := []string{
		"get 1",
		"get 2",
		"get 3",
		"get 4",
		"pow 1",
		"get 1",
		"get 2",
		"get 3",
		"get 4",
	}
	expect := []int{2, 0, 0, 1, 2, 1, 1, 0}
	runSample(t, n, P, a, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	a := []int{1}
	queries := []string{
		"pow 1",
		"get 1",
		"pow 1",
		"get 1",
	}
	expect := []int{0, 1}
	runSample(t, n, nil, a, queries, expect)
}
