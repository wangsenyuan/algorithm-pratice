package main

import "testing"

func runSample(t *testing.T, n int, correct string, chef string, W []int, expect int) {
	res := solve(n, correct, chef, W)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	correct := "ABCDE"
	chef := "EBCDA"
	W := []int{0, 10, 20, 30, 40, 50}
	expect := 30
	runSample(t, n, correct, chef, W, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	correct := "CHEF"
	chef := "QUIZ"
	W := []int{4, 3, 2, 1, 0}
	expect := 4
	runSample(t, n, correct, chef, W, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	correct := "ABBABAAB"
	chef := "ABABABAB"
	W := []int{100, 100, 100, 100, 100, 100, 100, 100, 100}
	expect := 100
	runSample(t, n, correct, chef, W, expect)
}
