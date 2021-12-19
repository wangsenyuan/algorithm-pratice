package main

import "testing"

func runSample(t *testing.T, G []string, expect int) {
	res := solve(len(G), G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	G := []string{
		"10",
		"00",
	}
	expect := 1
	runSample(t, G, expect)
}

func TestSample2(t *testing.T) {
	G := []string{
		"1111",
		"1101",
		"1011",
		"1111",
	}
	expect := 17
	runSample(t, G, expect)
}
