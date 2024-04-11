package main

import "testing"

func runSample(t *testing.T, m int, voters [][]int, expect int) {
	res := solve(m, voters)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	voters := [][]int{
		{1, 100},
	}
	expect := 0
	runSample(t, m, voters, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	voters := [][]int{
		{2, 100},
		{3, 200},
		{4, 300},
		{5, 400},
		{5, 900},
	}
	expect := 500
	runSample(t, m, voters, expect)
}

func TestSample3(t *testing.T) {
	m := 5
	voters := [][]int{
		{2, 100},
		{3, 200},
		{4, 300},
		{5, 800},
		{5, 900},
	}
	expect := 600
	runSample(t, m, voters, expect)
}
