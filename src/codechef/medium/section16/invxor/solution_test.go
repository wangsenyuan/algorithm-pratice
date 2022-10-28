package main

import "testing"

func runSample(t *testing.T, n string, x int, m int, expect int) {
	res := solve(n, m, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := "2"
	x := 2
	m := 7
	expect := 6
	runSample(t, n, x, m, expect)
}

func TestSample2(t *testing.T) {
	n := "4"
	x := 3
	m := 17
	expect := -1
	runSample(t, n, x, m, expect)
}

func TestSample3(t *testing.T) {
	n := "3"
	x := 5
	m := 13
	expect := 28
	runSample(t, n, x, m, expect)
}
