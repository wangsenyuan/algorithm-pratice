package main

import "testing"

func runSample(t *testing.T, n int64, labels []string, S []int, expect int) {
	res := solve(n, labels, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 5
	labels := []string{"A", "B"}
	S := []int{1, 2}
	expect := 16
	runSample(t, n, labels, S, expect)
}

func TestSample2(t *testing.T) {
	var n int64 = 6
	labels := []string{"A", "B", "C"}
	S := []int{1, 2, 3}
	expect := 113
	runSample(t, n, labels, S, expect)
}

func TestSample3(t *testing.T) {
	var n int64 = 8
	labels := []string{"A", "A", "B"}
	S := []int{2, 3, 2}
	expect := 128
	runSample(t, n, labels, S, expect)
}
