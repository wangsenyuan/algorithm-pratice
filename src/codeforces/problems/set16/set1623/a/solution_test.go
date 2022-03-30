package main

import "testing"

func runSample(t *testing.T, input []int, expect int) {
	res := solve(input[0], input[1], input[2], input[3], input[4], input[5])

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	input := []int{10, 10, 6, 1, 2, 8}
	expect := 7
	runSample(t, input, expect)
}

func TestSample2(t *testing.T) {
	input := []int{10, 10, 9, 9, 1, 1}
	expect := 10
	runSample(t, input, expect)
}

func TestSample3(t *testing.T) {
	input := []int{9, 8, 5, 6, 2, 1}
	expect := 9
	runSample(t, input, expect)
}

func TestSample4(t *testing.T) {
	input := []int{6, 9, 2, 2, 5, 8}
	expect := 3
	runSample(t, input, expect)
}
