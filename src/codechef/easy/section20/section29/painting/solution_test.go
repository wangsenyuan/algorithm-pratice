package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, commands []string, expect []int) {
	res := solve(len(commands), commands)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	commands := []string{
		"R 100",
		"U 10",
		"L 120",
		"D 10",
		"R 200",
	}
	expect := []int{100, 10, 120, 10, 99}
	runSample(t, commands, expect)
}
