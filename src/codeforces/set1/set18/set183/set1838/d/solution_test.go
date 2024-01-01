package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries []int, expect []bool) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "(())()()))"
	queries := []int{9, 7, 2, 6, 3, 6, 7, 4, 8}
	expect := []bool{true, true, false, false, true, false, true, false, false}
	runSample(t, s, queries, expect)
}
