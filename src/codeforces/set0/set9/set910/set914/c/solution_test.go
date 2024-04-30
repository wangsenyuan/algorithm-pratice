package main

import "testing"

func runSample(t *testing.T, num string, k int, expect int) {
	res := solve(num, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "110"
	k := 2
	expect := 3
	runSample(t, num, k, expect)
}

func TestSample2(t *testing.T) {
	num := "111111011"
	k := 2
	expect := 169
	runSample(t, num, k, expect)
}

func TestSample3(t *testing.T) {
	num := "1"
	k := 1
	expect := 0
	runSample(t, num, k, expect)
}

func TestSample4(t *testing.T) {
	num := "11111100010011110101100110100010001011100111001011001111101111110111001111011011110101100101001000111001000100000011011110110010001000111101001101001010100011"
	k := 1
	expect := 157
	runSample(t, num, k, expect)
}
