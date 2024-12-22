package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{12, 3, 45}
	expect := 12330
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12, 3}
	// 123
	// 132
	// 33
	// 1122
	expect := 255 + 33 + 1122
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{123, 456}
	expect := 1115598
	runSample(t, a, expect)
}
