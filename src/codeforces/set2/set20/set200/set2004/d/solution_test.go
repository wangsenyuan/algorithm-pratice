package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, cities []string, queries [][]int, expect []int) {
	res := solve(n, cities, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	cities := []string{
		"BR", "BR", "GY", "GR",
	}
	queries := [][]int{
		{1, 2},
		{3, 1},
		{4, 4},
		{1, 4},
		{4, 2},
	}
	expect := []int{1, 4, 0, 3, 2}
	runSample(t, n, cities, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	cities := []string{
		"BG", "RY",
	}
	queries := [][]int{
		{1, 2},
	}
	expect := []int{-1}
	runSample(t, n, cities, queries, expect)
}
