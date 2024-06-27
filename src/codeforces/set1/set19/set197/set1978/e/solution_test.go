package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a string, b string, queries [][]int, expect []int) {
	res := solve(a, b, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "1111"
	b := "0000"
	queries := [][]int{
		{1, 2},
		{2, 4},
	}
	expect := []int{2, 3}
	runSample(t, a, b, queries, expect)
}

func TestSample2(t *testing.T) {
	a := "1010"
	b := "1101"
	queries := [][]int{
		{1, 3},
		{1, 4},
	}
	expect := []int{2, 3}
	runSample(t, a, b, queries, expect)
}

func TestSample3(t *testing.T) {
	a := "010101"
	b := "011010"
	queries := [][]int{
		{2, 3},
		{1, 6},
		{2, 5},
		{4, 4},
		{3, 6},
	}
	expect := []int{1, 4, 3, 1, 2}
	runSample(t, a, b, queries, expect)
}

func TestSample4(t *testing.T) {
	a := "11110001010000111011010110101010001110001100100010"
	b := "10001100101100011110110010011110111010101100100101"
	queries := [][]int{
		{2, 32},
		{25, 41},
	}
	expect := []int{24, 15}
	runSample(t, a, b, queries, expect)
}
