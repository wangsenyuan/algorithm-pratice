package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, timetable []string, expect int) {
	res := solve(n, m, k, timetable)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 2, 5, 1
	timetable := []string{
		"01001",
		"10110",
	}
	expect := 5
	runSample(t, n, m, k, timetable, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 2, 5, 0
	timetable := []string{
		"01001",
		"10110",
	}
	expect := 8
	runSample(t, n, m, k, timetable, expect)
}
