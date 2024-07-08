package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "25"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "_00"
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "_XX"
	expect := 9
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0"
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "0_25"
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "_"
	expect := 1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "__"
	expect := 3
	runSample(t, s, expect)
}
