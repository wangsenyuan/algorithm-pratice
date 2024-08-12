package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0110"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "011011"
	// 110011
	// 001111
	// 011110
	// 110110
	// 111100
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "01010"
	// 110011
	// 001111
	// 011110
	// 110110
	// 111100
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "10001111110110111000"
	// 110011
	// 001111
	// 011110
	// 110110
	// 111100
	expect := 1287
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "00110110100110111101"
	// 110011
	// 001111
	// 011110
	// 110110
	// 111100
	expect := 1287
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "11101111011000100010"
	// 110011
	// 001111
	// 011110
	// 110110
	// 111100
	expect := 715
	runSample(t, s, expect)
}
