package main

import "testing"

func runSample(t *testing.T, k int, d int, expect int) {
	res := solve(k, d)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d, k := 1, 3
	expect := 1
	runSample(t, k, d, expect)
}


func TestSample2(t *testing.T) {
	d, k := 2, 6
	expect := 5
	runSample(t, k, d, expect)
}
