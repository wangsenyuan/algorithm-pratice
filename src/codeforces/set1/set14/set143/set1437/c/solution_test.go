package main

import "testing"

func runSample(t *testing.T, s []int, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []int{4, 2, 4, 4, 5, 2}
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []int{7, 7, 7, 7, 7, 7, 7}
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []int{5, 1, 2, 4, 3}
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []int{1, 4, 4, 4}
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := []int{21, 8, 1, 4, 1, 5, 21, 1, 8, 21, 11, 21, 11, 3, 12, 8, 19, 15, 9, 11, 13}
	expect := 21
	runSample(t, s, expect)
}
