package main

import "testing"

func runSample(t *testing.T, num int, m int, expect int) {
	res := solve(num, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num, m := 104, 2
	expect := 3
	runSample(t, num, m, expect)
}

func TestSample2(t *testing.T) {
	num, m := 223, 4
	expect := 1
	runSample(t, num, m, expect)
}

func TestSample3(t *testing.T) {
	num, m := 7067678, 8
	expect := 47
	runSample(t, num, m, expect)
}
