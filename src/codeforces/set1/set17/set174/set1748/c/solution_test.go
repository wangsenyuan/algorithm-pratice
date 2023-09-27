package main

import "testing"

func runSample(t *testing.T, v []int, expect int) {
	res := solve(v)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	v := []int{2, 0, 1, -1, 0}
	expect := 3
	runSample(t, v, expect)
}

func TestSample2(t *testing.T) {
	v := []int{1000000000, 1000000000, 0}
	expect := 1
	runSample(t, v, expect)
}

func TestSample3(t *testing.T) {
	v := []int{0, 0, 0, 0}
	expect := 4
	runSample(t, v, expect)
}

func TestSample5(t *testing.T) {
	v := []int{3, 0, 2, -10, 10, -30, 30, 0}
	expect := 4
	runSample(t, v, expect)
}

func TestSample6(t *testing.T) {
	v := []int{1, 0, 0, 1, -1, 0, 1, 0, -1}
	expect := 5
	runSample(t, v, expect)
}

func TestSample7(t *testing.T) {
	v := []int{-2, -2, 2}
	expect := 0
	runSample(t, v, expect)
}
