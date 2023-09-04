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
		"101",
		"111",
		"011",
		"110",
	}
	expect := 8
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"1110",
		"0111",
		"0111",
	}
	expect := 9
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"00",
		"00",
	}
	expect := 0
	runSample(t, mat, expect)
}

func TestSample4(t *testing.T) {
	mat := []string{
		"01",
		"10",
	}
	expect := 2
	runSample(t, mat, expect)
}
