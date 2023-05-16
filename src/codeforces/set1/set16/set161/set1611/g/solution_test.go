package main

import "testing"

func runSample(t *testing.T, field []string, expect int) {
	res := solve(field)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	field := []string{
		"00",
		"00",
	}
	var expect int
	runSample(t, field, expect)
}

func TestSample2(t *testing.T) {
	field := []string{
		"100",
		"000",
		"101",
	}
	var expect = 2
	runSample(t, field, expect)
}

func TestSample3(t *testing.T) {
	field := []string{
		"01000",
		"00001",
		"00010",
		"10000",
	}
	var expect = 2
	runSample(t, field, expect)
}

func TestSample4(t *testing.T) {
	field := []string{
		"111",
		"111",
		"111",
	}
	var expect = 4
	runSample(t, field, expect)
}

func TestSample5(t *testing.T) {
	field := []string{
		"11100",
		"11101",
	}
	var expect = 5
	runSample(t, field, expect)
}

func TestSample6(t *testing.T) {
	field := []string{
		"0001",
		"1101",
		"1001",
		"1010",
		"1110",
	}
	var expect = 4
	runSample(t, field, expect)
}

func TestSample7(t *testing.T) {
	field := []string{
		"10110",
		"11111",
		"01111",
	}
	var expect = 5
	runSample(t, field, expect)
}
