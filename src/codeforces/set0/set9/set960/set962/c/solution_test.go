package main

import "testing"

func runSample(t *testing.T, num int, expect int) {
	res := solve(num)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 8314
	expect := 2
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 625
	expect := 0
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	num := 333
	expect := -1
	runSample(t, num, expect)
}
