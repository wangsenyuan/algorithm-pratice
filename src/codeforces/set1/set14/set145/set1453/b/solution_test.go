package main

import "testing"

func runSample(t *testing.T, n int, arr []int, expect int64) {
	res := solve(n, arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	arr := []int{-1, 0, 2}
	var expect int64 = 1
	runSample(t, n, arr, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	arr := []int{99, 96, 97, 95}
	var expect int64 = 3
	runSample(t, n, arr, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	arr := []int{-3, -5, -2, 1}
	var expect int64 = 4
	runSample(t, n, arr, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	arr := []int{1, 4, 3, 2, 4, 1}
	var expect int64 = 6
	runSample(t, n, arr, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	arr := []int{5, 0, 0, 0, 5}
	var expect int64 = 5
	runSample(t, n, arr, expect)
}

func TestSample6(t *testing.T) {
	n := 9
	arr := []int{-367741579, 319422997,
		-415264583, -125558838, -300860379, 420848004, 294512916, -383235489, 425814447}
	var expect int64 = 2847372102
	runSample(t, n, arr, expect)
}
