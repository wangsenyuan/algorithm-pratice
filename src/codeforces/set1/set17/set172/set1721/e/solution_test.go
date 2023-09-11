package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, Q []string, expect [][]int) {
	res := solve(s, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aba"
	Q := []string{
		"caba",
		"aba",
		"bababa",
		"aaaa",
		"b",
		"forces",
	}
	expect := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3},
		{2, 3, 4, 5, 6, 7},
		{1, 1, 1, 1},
		{2},
		{0, 0, 0, 0, 0, 0},
	}
	runSample(t, s, Q, expect)
}

func TestSample2(t *testing.T) {
	s := "aacba"
	Q := []string{
		"aaca",
		"cbbb",
		"aab",
		"ccaca",
	}
	expect := [][]int{
		{2, 2, 3, 1},
		{0, 0, 0, 0},
		{2, 2, 0},
		{0, 0, 1, 0, 1},
	}
	runSample(t, s, Q, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaa"
	Q := []string{
		"aaaa",
	}
	expect := [][]int{
		{4, 5, 6, 7},
	}
	runSample(t, s, Q, expect)
}
