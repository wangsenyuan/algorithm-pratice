package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries []string, expect []int) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "47"
	queries := []string{
		"count",
		"switch 1 2",
		"count",
	}
	expect := []int{2, 1}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "747"
	queries := []string{
		"count",
		"switch 1 1",
		"count",
		"switch 1 3",
		"count",
	}
	expect := []int{2, 3, 2}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "7474747"
	queries := []string{
		"count",
		"count",
		"switch 1 7",
		"count",
		"switch 1 1",
		"switch 3 3",
		"switch 5 5",
		"count",
		"switch 1 7",
		"count",
	}
	expect := []int{4, 4, 4, 6, 7}
	runSample(t, s, queries, expect)
}
