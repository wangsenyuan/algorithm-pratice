package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, C int, Q []string, expect []int) {
	res := solve(C, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := 5
	Q := []string{
		"add 11101",
		"path 1 2",
		"add 01110",
		"path 3 2",
		"add 00100",
		"path 3 3",
		"add 10110",
		"remove",
		"path 1 2",
	}
	expect := []int{
		0, 1, 3, 2,
	}
	runSample(t, C, Q, expect)
}
