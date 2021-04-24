package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(len(s), s)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10101011011"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0000"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11111"
	expect := true
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := "1100"
	expect := false
	runSample(t, s, expect)
}
