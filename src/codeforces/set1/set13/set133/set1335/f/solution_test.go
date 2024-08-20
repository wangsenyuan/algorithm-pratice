package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, color []string, grid []string, expect []int) {
	res := solve(color, grid)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	color := []string{
		"01",
	}
	grid := []string{
		"RL",
	}
	expect := []int{2, 1}
	runSample(t, color, grid, expect)
}

func TestSample2(t *testing.T) {
	color := []string{
		"001",
		"101",
		"110",
	}
	grid := []string{
		"RLL",
		"DLD",
		"ULL",
	}
	expect := []int{4, 3}
	runSample(t, color, grid, expect)
}

func TestSample3(t *testing.T) {
	color := []string{
		"000",
		"000",
		"000",
	}
	grid := []string{
		"RRD",
		"RLD",
		"ULL",
	}
	expect := []int{2, 2}
	runSample(t, color, grid, expect)
}
