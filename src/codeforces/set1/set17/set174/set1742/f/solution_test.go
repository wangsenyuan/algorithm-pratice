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
		"2 1 aa",
		"1 2 a",
		"2 3 a",
		"1 2 b",
		"2 3 abca",
	}
	expect := []bool{true, false, true, false, true}
	runSample(t, queries, expect)
}

func TestSample2(t *testing.T) {
	queries := []string{
		"1 5 mihai",
		"2 2 buiucani",
	}
	expect := []bool{false, true}
	runSample(t, queries, expect)
}

func TestSample3(t *testing.T) {
	queries := []string{
		"1 5 b",
		"2 3 a",
		"2 4 paiu",
	}
	expect := []bool{false, false, true}
	runSample(t, queries, expect)
}
