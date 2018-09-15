package main

import "testing"

func runSample(t *testing.T, count []int, want []int, expect int) {
	res := solve(count, want)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", count, want, expect, res)
	}
}

func TestSample1(t *testing.T) {
	count := []int{2, 4, 6, 2, 0, 0, 6, 4, 2, 2}
	want := []int{1, 1, 0, 1, 0, 0, 1, 1, 0, 0}
	expect := 220
	runSample(t, count, want, expect)
}

func TestSample0(t *testing.T) {
	count := []int{2, 4, 6, 2, 0, 0, 6, 4, 2, 2}
	want := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 0}
	expect := 205
	runSample(t, count, want, expect)
}

func TestSample2(t *testing.T) {
	count := []int{2, 4, 6, 2, 0, 0, 6, 4, 2, 2}
	want := []int{0, 1, 1, 0, 0, 1, 1, 1, 0, 0}
	expect := 295
	runSample(t, count, want, expect)
}

func TestSample3(t *testing.T) {
	count := []int{2, 4, 6, 2, 0, 0, 6, 4, 2, 2}
	want := []int{1, 1, 0, 0, 0, 1, 1, 0, 0, 0}
	expect := 264
	runSample(t, count, want, expect)
}
