package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, C string, expect bool) {
	res := solve(n, E, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	C := "BRGB"
	E := [][]int{
		{3, 2},
		{1, 2},
		{4, 3},
	}
	expect := true
	runSample(t, n, E, C, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	C := "RGB"
	E := [][]int{
		{3, 2},
		{1, 2},
	}
	expect := true
	runSample(t, n, E, C, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	C := "BB"
	E := [][]int{
		{1, 2},
	}
	expect := false
	runSample(t, n, E, C, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	C := "BRGB"
	E := [][]int{
		{3, 2},
		{1, 2},
		{4, 2},
	}
	expect := false
	runSample(t, n, E, C, expect)
}
