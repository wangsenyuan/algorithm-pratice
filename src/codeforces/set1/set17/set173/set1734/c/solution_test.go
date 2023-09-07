package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(len(s), s)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111111"
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1101001"
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0000"
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0010"
	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "10010101"
	expect := 17
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "110011100101100"
	expect := 60
	runSample(t, s, expect)
}
