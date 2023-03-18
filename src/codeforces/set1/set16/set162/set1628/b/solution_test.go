package main

import "testing"

func runSample(t *testing.T, s []string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"zx",
		"ab",
		"cc",
		"zx",
		"ba",
	}
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"ab",
		"bad",
	}
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []string{
		"co",
		"def",
		"orc",
		"es",
	}
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []string{
		"a",
		"b",
		"c",
	}
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := []string{
		"ab",
		"cd",
		"cba",
	}
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := []string{
		"ab",
		"ab",
	}
	expect := false
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := []string{
		"abc",
		"ba",
	}
	expect := true
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := []string{
		"abc",
		"cba",
	}
	expect := true
	runSample(t, s, expect)
}
