package main

import "testing"

func runSample(t *testing.T, P, Q int, people []int, expect int) {
	res := solve(P, Q, people)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", P, Q, people, expect, res)
	}
}

func TestSample1(t *testing.T) {
	P, Q := 8, 1
	people := []int{3}
	expect := 7
	runSample(t, P, Q, people, expect)
}

func TestSample2(t *testing.T) {
	P, Q := 20, 3
	people := []int{3, 6, 14}
	expect := 35
	runSample(t, P, Q, people, expect)
}
