package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pos []int, dirs []int, questions [][]int, expect [][]int) {
	res := solve(pos, dirs, questions)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pos := []int{1, 24, 28, 34, 94}
	dirs := []int{0, 1, 0, 0, 0}
	questions := [][]int{
		{2, 11},
		{0, 9},
	}
	expect := [][]int{
		{23, 2},
		{-8, 0},
	}
	runSample(t, pos, dirs, questions, expect)
}
