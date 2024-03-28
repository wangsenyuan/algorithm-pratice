package main

import "testing"

func runSample(t *testing.T, m int, variables []string, expect []string) {
	a, b := solve(m, variables)

	if a != expect[0] || b != expect[1] {
		t.Fatalf("Sample expect %v, but got %s %s", expect, a, b)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	variables := []string{
		"a := 101",
		"b := 011",
		"c := ? XOR b",
	}
	expect := []string{
		"011",
		"100",
	}
	runSample(t, m, variables, expect)
}

func TestSample2(t *testing.T) {
	m := 1
	variables := []string{
		"a := 1",
		"bb := 0",
		"cx := ? OR a",
		"d := ? XOR ?",
		"e := d AND bb",
	}
	expect := []string{
		"0",
		"0",
	}
	runSample(t, m, variables, expect)
}
