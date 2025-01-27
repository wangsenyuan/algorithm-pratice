package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111010"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "10001"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1100010"
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "00100"
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1000101001"
	expect := 3
	runSample(t, s, expect)
}
