package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	expect := "101"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1010"
	expect := "11"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0000"
	expect := "0000"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0001"
	expect := "01"
	runSample(t, s, expect)
}
