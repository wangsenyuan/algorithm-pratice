package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, grid []string, expect []string) {
	res := solve(grid)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if len(expect) == 0 {
		return
	}
	sort.Strings(expect)
	sort.Strings(res)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		`.//.`,
		`.\\.`,
		`.\/.`,
		`....`,
	}
	expect := []string{"N3", "W2"}
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		`./..\.`,
		`.\...\`,
		`./../\`,
		`......`,
	}
	expect := []string{"E3", "S2"}
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		`....`,
		`./\.`,
		`.\/.`,
		`....`,
	}
	runSample(t, grid, nil)
}

func TestSample4(t *testing.T) {
	grid := []string{
		`/\.`,
		`\\\`,
		`\/\`,
	}
	expect := []string{"W2", "E3"}
	runSample(t, grid, expect)
}
