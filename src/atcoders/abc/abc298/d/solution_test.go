package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cmds []string, expect []int) {
	res := solve(cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cmds := []string{
		"3",
		"1 2",
		"3",
	}
	expect := []int{1, 12}
	runSample(t, cmds, expect)
}

func TestSample2(t *testing.T) {
	cmds := []string{
		"1 5",
		"2",
		"3",
	}
	expect := []int{5}
	runSample(t, cmds, expect)
}

func TestSample3(t *testing.T) {
	cmds := []string{
		"1 9",
		"1 9",
		"1 8",
		"1 2",
		"1 4",
		"1 4",
		"1 3",
		"1 5",
		"1 3",
		"2",
		"3",
	}
	expect := []int{0}
	runSample(t, cmds, expect)
}
