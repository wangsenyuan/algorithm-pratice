package main

import "testing"

func runSample(t *testing.T, n int, P1 int, P2 int, P3 int, T1 int, T2 int, events [][]int, expect int) {
	res := solve(n, P1, P2, P3, T1, T2, events)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, P1, P2, P3, T1, T2 := 1, 3, 2, 1, 5, 10
	events := [][]int{
		{0, 10},
	}
	expect := 30
	runSample(t, n, P1, P2, P3, T1, T2, events, expect)
}

func TestSample2(t *testing.T) {
	n, P1, P2, P3, T1, T2 := 2, 8, 4, 2, 5, 10
	events := [][]int{
		{20, 30},
		{50, 100},
	}
	expect := 570
	runSample(t, n, P1, P2, P3, T1, T2, events, expect)
}
