package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []bool) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "+-+---+-"
	queries := [][]int{
		{2, 1},
		{10, 3},
		{7, 9},
		{10, 10},
		{5, 3},
	}
	expect := []bool{true, false, false, false, true}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "+-++--"
	queries := [][]int{
		{9, 7},
		{1, 1},
	}
	expect := []bool{true, true}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "+-----+--+--------+-"
	queries := [][]int{
		{1000000000, 99999997},
		{250000000, 1000000000},
	}
	expect := []bool{false, true}
	runSample(t, s, queries, expect)
}
