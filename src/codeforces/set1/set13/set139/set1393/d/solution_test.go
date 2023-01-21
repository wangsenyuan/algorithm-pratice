package main

import "testing"

func runSample(t *testing.T, m int, n int, G []string, expect int) {
	res := solve(m, n, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 3, 3
	G := []string{
		"aaa",
		"aaa",
		"aaa",
	}
	expect := 10
	runSample(t, m, n, G, expect)
}

func TestSample2(t *testing.T) {
	m, n := 3, 4
	G := []string{
		"abab",
		"baba",
		"abab",
	}
	expect := 12
	runSample(t, m, n, G, expect)
}

func TestSample3(t *testing.T) {
	m, n := 5, 5
	G := []string{
		"zbacg",
		"baaac",
		"aaaaa",
		"eaaad",
		"weadd",
	}
	expect := 31
	runSample(t, m, n, G, expect)
}
