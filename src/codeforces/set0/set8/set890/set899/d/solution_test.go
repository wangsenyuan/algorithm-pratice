package main

import "testing"

func runSample(t *testing.T, num int, expect int) {
	res := solve(num)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 7
	expect := 3
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 14
	expect := 9
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	num := 50
	expect := 1
	runSample(t, num, expect)
}
