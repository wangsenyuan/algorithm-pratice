package main

import "testing"

func runSample(t *testing.T, p []int, s string, expect int) {
	res := solve(p, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1, 2, 3}
	s := "L??"
	expect := 2
	runSample(t, p, s, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 3, 2}
	s := "R?L"
	expect := 0
	runSample(t, p, s, expect)
}

func TestSample3(t *testing.T) {
	p := []int{6, 2, 9, 3, 1, 4, 11, 5, 12, 10, 7, 8}
	s := "????????????"
	expect := 160
	runSample(t, p, s, expect)
}

func TestSample4(t *testing.T) {
	p := []int{1, 2}
	s := "LR"
	expect := 1
	runSample(t, p, s, expect)
}
