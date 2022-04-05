package main

import "testing"

func runSample(t *testing.T, H []int, expect int) {
	res := solve(H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{2, 1}
	expect := 3
	runSample(t, H, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1}
	expect := 2
	runSample(t, H, expect)
}

func TestSample3(t *testing.T) {
	H := []int{1, 2, 3, 4, 5}
	expect := 32
	runSample(t, H, expect)
}

func TestSample4(t *testing.T) {
	H := []int{1, 3, 4, 2}
	expect := 7
	runSample(t, H, expect)
}

func TestSample5(t *testing.T) {
	H := []int{6, 1, 7, 3, 5, 4, 2}
	expect := 11
	runSample(t, H, expect)
}
