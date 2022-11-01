package main

import "testing"

func runSample(t *testing.T, l, r int, expect int64) {
	res := solve(l, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r := 1, 4
	expect := int64(3)
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 3, 5
	expect := int64(1)
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 8, 86
	expect := int64(78975)
	runSample(t, l, r, expect)
}

func TestSample4(t *testing.T) {
	l, r := 68, 86
	expect := int64(969)
	runSample(t, l, r, expect)
}

func TestSample5(t *testing.T) {
	l, r := 6, 86868
	expect := int64(109229059713337)
	runSample(t, l, r, expect)
}
