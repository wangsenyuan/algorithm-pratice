package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, rounds []int, expect []string) {
	res := solve(n, edges, rounds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{2, 3},
		{3, 1},
	}
	rounds := []int{3}
	expect := []string{Ron}

	runSample(t, n, edges, rounds, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	rounds := []int{5}
	expect := []string{Hermione}

	runSample(t, n, edges, rounds, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	rounds := []int{1}
	expect := []string{Ron}

	runSample(t, n, edges, rounds, expect)
}
