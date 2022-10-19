package main

import "testing"

func runSample(t *testing.T, n int, s []string, blocks [][]int, expect bool) {
	res := solve(n, s, blocks)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	s := []string{
		"010",
		"111",
		"010",
	}
	blocks := [][]int{
		{2, 2},
	}
	expect := false
	runSample(t, n, s, blocks, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	s := []string{
		"1110",
		"1110",
		"1001",
		"1101",
	}
	blocks := [][]int{
		{2, 2},
		{4, 3},
	}
	expect := true
	runSample(t, n, s, blocks, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	s := []string{
		"10100",
		"01100",
		"10011",
		"00010",
		"00100",
	}
	blocks := [][]int{
		{3, 5},
		{5, 1},
		{5, 3},
	}
	expect := true
	runSample(t, n, s, blocks, expect)
}
