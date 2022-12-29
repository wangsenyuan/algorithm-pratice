package main

import "testing"

func runSample(t *testing.T, k int, s string, expect int) {
	res := solve(len(s), k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "000"
	k := 3
	expect := 1
	runSample(t, k, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0110"
	k := 2
	expect := 0
	runSample(t, k, s, expect)
}


func TestSample3(t *testing.T) {
	s := "10101"
	k := 3
	expect := 2
	runSample(t, k, s, expect)
}

func TestSample4(t *testing.T) {
	s := "000"
	k := 1
	expect := 1
	runSample(t, k, s, expect)
}

func TestSample5(t *testing.T) {
	s := "010"
	k := 3
	expect := 3
	runSample(t, k, s, expect)
}
