package main

import "testing"

func runSample(t *testing.T, time int, field []string, trees [][]int, expect int) {
	res := solve(time, field, trees)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	T := 30
	fields := []string{
		"...",
		".#.",
		".#S",
	}
	trees := [][]int{
		{3, 1, 0, 20, 4},
	}
	expect := 56
	runSample(t, T, fields, trees, expect)
}

func TestSample2(t *testing.T) {
	T := 20
	fields := []string{
		"SS..",
	}
	trees := [][]int{
		{1, 3, 0, 4, 20},
		{1, 4, 0, 4, 10},
	}
	expect := 60
	runSample(t, T, fields, trees, expect)
}
