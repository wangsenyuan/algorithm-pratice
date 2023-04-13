package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q []string, expect []int64) {
	res := solve(n, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	Q := []string{
		"Color 2 4 2",
		"Add 2 2",
		"Query 3",
		"Color 4 5 3",
		"Color 2 2 3",
		"Add 3 3",
		"Query 2",
		"Query 5",
	}
	expect := []int64{2, 5, 3}
	runSample(t, n, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	Q := []string{
		"Add 1 7",
		"Query 1",
		"Add 2 4",
		"Query 2",
		"Color 1 1 1",
		"Add 1 1",
		"Query 2",
	}
	expect := []int64{7, 7, 8}
	runSample(t, n, Q, expect)
}
