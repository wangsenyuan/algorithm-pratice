package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, k int, queries []string, expect []int) {
	res := solve(s, k, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	s := "0110?11"
	queries := []string{
		"5 1",
		"6 ?",
		"7 ?",
		"1 ?",
		"5 ?",
		"1 1",
	}
	expect := []int{1, 2, 3, 3, 5, 4}
	runSample(t, s, k, queries, expect)
}
