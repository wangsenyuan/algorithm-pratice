package main

import "testing"

func runSample(t *testing.T, mat []string, expect int) {
	res := solve(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"1000",
		"0000",
		"0110",
		"0000",
		"0001",
	}
	expect := 12
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"001010001",
		"101110100",
		"000010011",
		"100000001",
		"101010101",
		"110001111",
		"000001111",
		"111100000",
		"000110000",
	}
	expect := 5
	runSample(t, mat, expect)
}
