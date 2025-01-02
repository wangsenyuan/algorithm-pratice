package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "010101"
	expect := 3
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := "11001100"
	expect := 0
	runSample(t, s, expect)
}
