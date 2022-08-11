package main

import "testing"

func runSample(t *testing.T, first []int, second []int, expect int) {
	res := solve(first, second)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	first := []int{1, 1, 10, 10}
	second := []int{11, 11, 20, 20}
	expect := 162
	runSample(t, first, second, expect)
}

func TestSample2(t *testing.T) {
	first := []int{1, 1, 20, 20}
	second := []int{11, 11, 30, 30}
	expect := 641
	runSample(t, first, second, expect)
}
