package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []string, S string, C []int, expect []bool) {
	res := solve(P, S, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []string{
		"abc",
		"abcd",
		"ab",
		"abcde",
		"a",
		"abcd",
	}
	S := "abcd"
	C := []int{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	expect := []bool{false, false, false, true, false, true, true}
	runSample(t, P, S, C, expect)
}
