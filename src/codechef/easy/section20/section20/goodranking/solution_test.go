package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S []string, expect []int) {
	res := solve(n, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	S := []string{
		"0101011",
		"0001100",
		"1100011",
		"0010100",
		"1010011",
		"0101000",
		"0101010",
	}
	expect := []int{5, 3, 1, 7, 6, 2, 4}
	runSample(t, n, S, expect)
}
