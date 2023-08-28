package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid []string, expect []string) {
	res := solve(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		".*.*....*.",
		".*.......*",
		"...o....o.",
		".*.*....*.",
		"..........",
		".o......o*",
	}
	expect := []string{
		"..........",
		"...*....*.",
		".*.o....o.",
		".*........",
		".*......**",
		".o.*....o*",
	}
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"...***ooo",
		".*o.*o.*o",
	}
	expect := []string{
		"....**ooo",
		".*o**o.*o",
	}
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"*****",
		"*....",
		"*****",
		"....*",
		"*****",
	}
	expect := []string{
		".....",
		"*...*",
		"*****",
		"*****",
		"*****",
	}
	runSample(t, grid, expect)
}
