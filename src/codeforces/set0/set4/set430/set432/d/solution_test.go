package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	res := solve(s)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ABACABA"
	expect := [][]int{
		{1, 4},
		{3, 2},
		{7, 1},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "AAA"
	expect := [][]int{
		{1, 3},
		{2, 2},
		{3, 1},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "AXAXA"
	expect := [][]int{
		{1, 3},
		{3, 2},
		{5, 1},
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "AAAAAAAAAAAAAAAAXAAAAAAAAAAAAAAAAAAAAAAA"
	expect := [][]int{
		{1, 39},
		{2, 37},
		{3, 35},
		{4, 33},
		{5, 31},
		{6, 29},
		{7, 27},
		{8, 25},
		{9, 23},
		{10, 21},
		{11, 19},
		{12, 17},
		{13, 15},
		{14, 13},
		{15, 11},
		{16, 9},
		{40, 1},
	}
	runSample(t, s, expect)
}
