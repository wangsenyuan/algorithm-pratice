package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries []string, expect []bool) {
	res := solve(queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := []string{
		"+ 1 -1",
		"? 1 1 2",
		"? 1 2 1",
		"+ 1 1",
		"? 1 3 -1",
		"? 1 1 1",
		"? 1 3 2",
		"? 1 1 0",
	}
	expect := []bool{false, true, false, true, true, true}
	runSample(t, queries, expect)
}

func TestSample2(t *testing.T) {
	queries := []string{
		"+ 1 -1",
		"+ 1 -1",
		"+ 3 1",
		"+ 3 -1",
		"+ 3 1",
		"? 1 6 -1",
		"? 1 6 2",
		"? 1 2 0",
		"? 1 5 -2",
		"? 1 4 3",
	}
	expect := []bool{true, false, true, true, false}
	runSample(t, queries, expect)
}
