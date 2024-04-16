package main

import "testing"

func runSample(t *testing.T, persons [][]int, k int, expect bool) {
	res := solve(persons, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 12
	persons := [][]int{
		{5, 7, 10},
		{10, 16, 20},
	}
	expect := false
	runSample(t, persons, k, expect)
}
