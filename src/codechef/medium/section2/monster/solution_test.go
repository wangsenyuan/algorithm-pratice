package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, monsters []int, q int, actions [][]int, expect []int) {
	res := solve(n, monsters, q, actions)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, monsters, q, actions, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	monsters := []int{1, 2, 3, 4, 5}
	q := 5
	actions := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}
	expect := []int{4, 4, 2, 2, 1}
	runSample(t, n, monsters, q, actions, expect)
}
