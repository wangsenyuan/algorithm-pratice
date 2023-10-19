package main

import "testing"

func runSample(t *testing.T, walls []string, expect bool) {
	res := solve(walls)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	walls := []string{
		"WBB",
		"BBW",
	}
	expect := true
	runSample(t, walls, expect)
}

func TestSample2(t *testing.T) {
	walls := []string{
		"B",
		"B",
	}
	expect := true
	runSample(t, walls, expect)
}

func TestSample3(t *testing.T) {
	walls := []string{
		"BWBWB",
		"BBBBB",
	}
	expect := false
	runSample(t, walls, expect)
}

func TestSample4(t *testing.T) {
	walls := []string{
		"BW",
		"WB",
	}
	expect := false
	runSample(t, walls, expect)
}

func TestSample5(t *testing.T) {
	walls := []string{
		"BBBBW",
		"BWBBB",
	}
	expect := false
	runSample(t, walls, expect)
}

func TestSample6(t *testing.T) {
	walls := []string{
		"BWBBWB",
		"BBBBBB",
	}
	expect := true
	runSample(t, walls, expect)
}
