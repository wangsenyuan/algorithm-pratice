package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, edges [][]int, tickets [][]int, Q int, friends []int, expect []int) {
	res := solve(n, m, edges, tickets, Q, friends)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v %d %v, expect %v, but got %v", n, m, edges, tickets, Q, friends, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 7, 7
	edges := [][]int{
		{3, 1},
		{2, 1},
		{7, 6},
		{6, 3},
		{5, 3},
		{4, 3},
	}
	tickets := [][]int{
		{7, 2, 3},
		{7, 1, 1},
		{2, 3, 5},
		{3, 6, 2},
		{4, 2, 4},
		{5, 3, 10},
		{6, 1, 20},
	}
	Q := 3
	friends := []int{5, 6, 7}
	expect := []int{10, 22, 5}

	runSample(t, n, m, edges, tickets, Q, friends, expect)
}
