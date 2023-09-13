package main

import "testing"

func runSample(t *testing.T, m int, D []int, expect int) {
	res := solve(m, D)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	D := []int{10, 10, 6, 14}
	expect := 10
	runSample(t, m, D, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	D := []int{10}
	expect := 2
	runSample(t, m, D, expect)
}

func TestSample3(t *testing.T) {
	m := 2
	D := []int{1, 1000000000, 1000000000, 1000000000, 999999999}
	expect := 8
	runSample(t, m, D, expect)
}
