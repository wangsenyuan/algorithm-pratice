package main

import "testing"

func runSample(t *testing.T, n int, D []int64, expect bool) {
	res := solve(n, D)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	D := []int64{8, 12, 8, 12}
	expect := true
	runSample(t, n, D, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	D := []int64{240, 154, 210, 162, 174, 154, 186, 240, 174, 186, 162, 210}
	expect := true
	runSample(t, n, D, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	D := []int64{7, 7, 9, 11}
	expect := false
	runSample(t, n, D, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	D := []int64{7, 11, 7, 11}
	expect := false
	runSample(t, n, D, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	D := []int64{40, 56, 48, 40, 80, 56, 80, 48}
	expect := false
	runSample(t, n, D, expect)
}
