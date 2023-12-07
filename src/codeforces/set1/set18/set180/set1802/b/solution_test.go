package main

import "testing"

func runSample(t *testing.T, b []int, expect int) {
	res := solve(b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	b := []int{1, 1, 1}
	expect := 3
	runSample(t, b, expect)
}

func TestSample2(t *testing.T) {
	b := []int{2, 2, 2}
	expect := 0
	runSample(t, b, expect)
}

func TestSample3(t *testing.T) {
	b := []int{1, 1, 1, 2, 1}
	expect := 3
	runSample(t, b, expect)
}

func TestSample4(t *testing.T) {
	b := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	expect := 4
	runSample(t, b, expect)
}

func TestSample5(t *testing.T) {
	b := []int{1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1}
	expect := 12
	runSample(t, b, expect)
}

func TestSample6(t *testing.T) {
	b := []int{2, 1, 1, 2, 1, 1, 2, 1, 2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1, 2}
	expect := 9
	runSample(t, b, expect)
}
