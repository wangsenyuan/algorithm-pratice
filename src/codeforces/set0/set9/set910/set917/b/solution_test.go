package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, letters string, expect []string) {
	res := solve(n, edges, []byte(letters))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample  expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 4},
	}
	letters := "bacb"
	expect := []string{
		"BAAA",
		"ABAA",
		"BBBA",
		"BBBB",
	}
	runSample(t, n, edges, letters, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{5, 3},
		{1, 2},
		{3, 1},
		{3, 2},
		{5, 1},
		{4, 3},
		{5, 4},
		{5, 2},
	}
	letters := "hccrrzrh"
	expect := []string{
		"BABBB",
		"BBBBB",
		"AABBB",
		"AAABA",
		"AAAAB",
	}
	runSample(t, n, edges, letters, expect)
}
