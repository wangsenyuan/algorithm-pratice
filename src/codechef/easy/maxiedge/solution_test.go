package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, grid []string, expect []string) {
	res := solve(n, grid)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v, expect %v, but got %v", n, grid, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 2
	grid := []string{
		"ab",
		"ba",
	}
	expect := []string{
		"ab",
		"ab",
	}
	runSample(t, n, grid, expect)
}
