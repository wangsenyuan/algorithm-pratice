package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, edges [][]int, expect []int) {
	res := solve(len(a), edges, a)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 1, 0}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := []int{8, 6, 12, 10}
	runSample(t, a, edges, expect)
}

func TestSample2(t *testing.T) {
	a := []int{100}
	edges := [][]int{}
	expect := []int{0}
	runSample(t, a, edges, expect)
}
