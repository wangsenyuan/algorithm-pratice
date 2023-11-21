package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "100"
	expect := 1000000000001
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0"
	expect := 0
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := "0101"
	expect := 1000000000000
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "00101101"
	expect := 2000000000001
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1001101"
	expect := 2000000000002
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "1111"
	expect := 0
	runSample(t, s, expect)
}
