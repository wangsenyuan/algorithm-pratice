package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := calculate(s)
	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "161+6033/2/69-11*2-12*6+51*40"
	runSample(t, s, 2150)
}

func TestSample2(t *testing.T) {
	s := "3+2*2"
	runSample(t, s, 7)
}

func TestSample3(t *testing.T) {
	s := " 3/2 "
	runSample(t, s, 1)
}

func TestSample4(t *testing.T) {
	s := " 3+5 / 2 "
	runSample(t, s, 5)
}
