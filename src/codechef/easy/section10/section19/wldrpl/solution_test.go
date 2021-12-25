package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S string, Q [][]int, expect []int) {
	res := solve(len(S), S, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "(?+?)"
	Q := [][]int{
		{2, 2},
		{1, 5},
	}
	expect := []int{1, 2}
	runSample(t, S, Q, expect)
}

func TestSample2(t *testing.T) {
	S := "(?-(?-?))"
	Q := [][]int{
		{1, 9},
		{4, 8},
	}
	expect := []int{2, 1}
	runSample(t, S, Q, expect)
}

func TestSample3(t *testing.T) {
	S := "(?-(?+?))"
	Q := [][]int{
		{1, 9},
		{4, 8},
	}
	expect := []int{1, 2}
	runSample(t, S, Q, expect)
}