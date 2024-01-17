package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r := 1, 20
	expect := 21
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 5, 199
	expect := 233
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 0, 1000000000
	expect := 2553052375
	runSample(t, l, r, expect)
}

func TestSample4(t *testing.T) {
	l, r := 0, 0
	expect := 1
	runSample(t, l, r, expect)
}

func TestSample5(t *testing.T) {
	var l, r int = 0, 1e18
	expect := 4076182193177635939
	runSample(t, l, r, expect)
}
