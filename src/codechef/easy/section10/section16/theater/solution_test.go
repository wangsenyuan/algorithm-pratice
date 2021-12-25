package main

import "testing"

func runSample(t *testing.T, n int, requests []string, expect int) {
	res := solve(n, requests)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, requests, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 12
	requests := []string{
		"A 3",
		"B 12",
		"C 6",
		"A 9",
		"B 12",
		"C 12",
		"D 3",
		"B 9",
		"D 3",
		"B 12",
		"B 9",
		"C 6",
	}
	expect := 575
	runSample(t, n, requests, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	requests := []string{
		"A 9",
		"A 9",
		"B 6",
		"C 3",
		"D 12",
		"A 9",
		"B 6",
	}
	expect := 525
	runSample(t, n, requests, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	requests := []string{
		"A 9",
		"B 6",
	}
	expect := -25
	runSample(t, n, requests, expect)
}

func TestSample4(t *testing.T) {
	n := 1
	requests := []string{
		"D 12",
	}
	expect := -200
	runSample(t, n, requests, expect)
}

func TestSample5(t *testing.T) {
	n := 0
	requests := []string{}
	expect := -400
	runSample(t, n, requests, expect)
}
