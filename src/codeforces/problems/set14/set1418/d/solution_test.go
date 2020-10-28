package main

import (
	"bytes"
	"fmt"
	"testing"
)

func runSample(t *testing.T, n int, P []int, q int, queries [][]int, expect []int) {
	res := solve(n, P, q, queries)

	var buf bytes.Buffer
	for i := 0; i < len(expect); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", expect[i]))
	}
	str := buf.String()
	if res != str {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 5, 6
	P := []int{1, 2, 6, 8, 10}
	queries := [][]int{
		{1, 4},
		{1, 9},
		{0, 6},
		{0, 10},
		{1, 100},
		{1, 50},
	}
	expect := []int{5, 7, 7, 5, 4, 8, 49}
	runSample(t, n, P, q, queries, expect)
}

func TestSample2(t *testing.T) {
	n, q := 5, 8
	P := []int{5, 1, 2, 4, 3}
	queries := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{1, 1000000000},
		{1, 1},
		{1, 500000000},
	}
	expect := []int{3, 2, 1, 0, 0, 0, 0, 0, 499999999}
	runSample(t, n, P, q, queries, expect)
}
