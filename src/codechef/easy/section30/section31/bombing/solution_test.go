package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, cmds []string, expect []int) {
	res := solve(n, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	cmds := []string{
		"P 1 4",
		"P 3 5",
		"B 3",
		"M 2 1",
		"B 3",
	}
	expect := []int{2, 1}
	runSample(t, n, cmds, expect)
}
