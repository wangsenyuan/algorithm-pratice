package main

import "testing"

func runSample(t *testing.T, first int, expect int) {
	res := solve(first)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	num := 11
	expect := 4
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 239
	expect := 28
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	num := 401
	expect := 61
	runSample(t, num, expect)
}

func TestSample4(t *testing.T) {
	num := 9001
	expect := 2834
	runSample(t, num, expect)
}
