package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r := 1, 2
	expect := 1
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 2, 8
	expect := 3
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 100000, 200000
	expect := 31072
	runSample(t, l, r, expect)
}
