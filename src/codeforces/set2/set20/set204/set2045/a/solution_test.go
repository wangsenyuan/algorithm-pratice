package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ICPCJAKARTA"
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "NGENG"
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "YYY"
	expect := 3
	runSample(t, s, expect)
}
func TestSample4(t *testing.T) {
	s := "DANGAN"
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "AEIOUY"
	expect := 0
	runSample(t, s, expect)
}
