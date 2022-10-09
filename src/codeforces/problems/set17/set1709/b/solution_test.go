package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, quests [][]int, expect []int64) {
	res := solve(A, quests)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 8, 9, 6, 8, 12, 7}
	quests := [][]int{
		{1, 2},
		{1, 7},
		{4, 6},
		{7, 1},
		{3, 5},
		{4, 2},
	}
	expect := []int64{2, 10, 0, 7, 3, 1}
	runSample(t, A, quests, expect)
}
