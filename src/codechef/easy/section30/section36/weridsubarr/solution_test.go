package main

import "testing"

func runSample(t *testing.T, P []int, expect int64) {
	res := solve(P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{2, 3, 1}
	expect := 5
	runSample(t, P, int64(expect))
}

func TestSample2(t *testing.T) {
	P := []int{2, 1, 3}
	expect := 6
	runSample(t, P, int64(expect))
}
func TestSample3(t *testing.T) {
	P := []int{1, 2, 3, 4}
	expect := 10
	runSample(t, P, int64(expect))
}
