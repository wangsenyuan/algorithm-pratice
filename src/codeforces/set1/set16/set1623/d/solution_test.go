package main

import "testing"

func runSample(t *testing.T, input []int, expect int) {
	res := solve(input[0], input[1], input[2], input[3], input[4], input[5], input[6])

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	input := []int{2, 2, 1, 1, 2, 1, 25}
	expect := 3
	runSample(t, input, expect)
}

func TestSample2(t *testing.T) {
	input := []int{3, 3, 1, 2, 2, 2, 25}
	expect := 3
	runSample(t, input, expect)
}

func TestSample3(t *testing.T) {
	input := []int{10, 10, 1, 1, 10, 10, 75}
	expect := 15
	runSample(t, input, expect)
}

func TestSample4(t *testing.T) {
	input := []int{10, 10, 10, 10, 1, 1, 75}
	expect := 15
	runSample(t, input, expect)
}

func TestSample5(t *testing.T) {
	input := []int{5, 5, 1, 3, 2, 2, 10}
	expect := 332103349
	runSample(t, input, expect)
}

func TestSample6(t *testing.T) {
	input := []int{97, 98, 3, 5, 41, 43, 50}
	expect := 99224487
	runSample(t, input, expect)
}
