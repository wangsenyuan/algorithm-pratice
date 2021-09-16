package main

import "testing"

func runSample(t *testing.T, num int, l, r int, expect int) {
	res := solve(num, l, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 216
	l := 2
	r := 7
	expect := 6
	runSample(t, num, l, r, expect)
}

func TestSample2(t *testing.T) {
	num := 256
	l := 2
	r := 4
	expect := 2
	runSample(t, num, l, r, expect)
}

func TestSample3(t *testing.T) {
	num := 31
	l := 3
	r := 5
	expect := 3
	runSample(t, num, l, r, expect)
}
