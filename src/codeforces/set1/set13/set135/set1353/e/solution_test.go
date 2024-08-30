package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "010001010"
	k := 2
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "111100000"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "1111111"
	k := 4
	expect := 5
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "1001110101"
	k := 3
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "0"
	k := 1
	expect := 0
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	s := "1"
	k := 1
	expect := 0
	runSample(t, s, k, expect)
}