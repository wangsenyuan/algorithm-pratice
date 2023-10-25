package main

import (
	"testing"
)

func runSample(t *testing.T, contests [][]int, expect string) {
	res := solve(contests)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	contests := [][]int{
		{-7, 1},
		{5, 2},
		{8, 2},
	}
	expect := "1907"
	runSample(t, contests, expect)
}

func TestSample2(t *testing.T) {
	contests := [][]int{
		{57, 1},
		{22, 2},
	}
	expect := "Impossible"
	runSample(t, contests, expect)
}

func TestSample3(t *testing.T) {
	contests := [][]int{
		{-5, 1},
	}
	expect := "Infinity"
	runSample(t, contests, expect)
}

func TestSample4(t *testing.T) {
	contests := [][]int{
		{27, 2},
		{13, 1},
		{-50, 1},
		{8, 2},
	}
	expect := "1897"
	runSample(t, contests, expect)
}
