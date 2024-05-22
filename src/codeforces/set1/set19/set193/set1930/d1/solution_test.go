package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "10"
	expect := 2
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := "11110110000000111111"
	expect := 346
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "11"
	expect := 3
	runSample(t, s, expect)
}