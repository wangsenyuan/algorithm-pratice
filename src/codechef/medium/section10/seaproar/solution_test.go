package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1101100100101111010011010101110100001000000101001110101011010101010"
	expect := "LCG"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "000101101110101101110110010111000000011001101110101"
	expect := "LCG"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11010100010110001101010110111000010001110010010011011110010010110000001100110"
	expect := "LCG"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "01011010100111100111101001010010100100111000111110"
	expect := "Xorshift"
	runSample(t, s, expect)
}
