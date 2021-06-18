package main

import "testing"

func runSample(t *testing.T, n int, S string, P []int, expect int64) {
	res := solve(n, S, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	S := "01010"
	P := []int{1, 1, 3, 3}
	var expect int64 = 2
	runSample(t, n, S, P, expect)
}


func TestSample2(t *testing.T) {
	n := 5
	S := "10101"
	P := []int{1, 1, 3, 3}
	var expect int64
	runSample(t, n, S, P, expect)
}
