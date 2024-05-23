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
		"WWWWWWW",
		"WWWWBBB",
		"WWWWWBW",
		"WWBBBBB",
		"WWWBWWW",
		"WWBBBWW",
		"WWWWWWW",
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"WWWWWWW",
		"WWWWWWW",
		"WBBBBBW",
		"WBBBBBW",
		"WBBBBBW",
		"WWWWWWW",
		"WWWWWWW",
	}
	expect := 2
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"WWWWWWW",
		"WWWWWWW",
		"WWWWWWW",
		"WWWWWWW",
		"WWWWWWW",
		"WWWWWWW",
		"WWWWWWW",
	}
	expect := 0
	runSample(t, mat, expect)
}

func TestSample4(t *testing.T) {
	mat := []string{
		"WBBBBBW",
		"BBBBBBB",
		"BBBBBBB",
		"WWWWWWW",
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
	}
	expect := 5
	runSample(t, mat, expect)
}

func TestSample5(t *testing.T) {
	mat := []string{
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
		"BBBBBBB",
	}
	expect := 8
	runSample(t, mat, expect)
}
