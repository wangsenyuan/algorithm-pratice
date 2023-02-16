package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, G []string, expect []int) {
	res := solve(n, G)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	G := []string{
		"000",
		"122",
		"001",
	}

	expect := []int{4, 4, 1, 0, 0, 0, 0, 0, 0, 0}
	runSample(t, n, G, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	G := []string{
		"57",
		"75",
	}

	expect := []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 0}
	runSample(t, n, G, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	G := []string{
		"0123",
		"4012",
		"3401",
		"2340",
	}

	expect := []int{9, 6, 9, 9, 6, 0, 0, 0, 0, 0}
	runSample(t, n, G, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	G := []string{
		"42987101",
		"98289412",
		"38949562",
		"87599023",
		"92834718",
		"83917348",
		"19823743",
		"38947912",
	}

	expect := []int{18, 49, 49, 49, 49, 15, 0, 30, 42, 42}
	runSample(t, n, G, expect)
}
