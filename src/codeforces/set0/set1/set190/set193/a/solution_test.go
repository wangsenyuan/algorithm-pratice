package main

import "testing"

func runSample(t *testing.T, G []string, expect int) {
	res := solve(G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	G := []string{
		"####",
		"#..#",
		"#..#",
		"#..#",
		"####",
	}
	expect := 2
	runSample(t, G, expect)
}

func TestSample2(t *testing.T) {
	G := []string{
		"#####",
		"#...#",
		"#####",
		"#...#",
		"#####",
	}
	expect := 2
	runSample(t, G, expect)
}

func TestSample3(t *testing.T) {
	G := []string{
		".########.",
	}
	expect := 1
	runSample(t, G, expect)
}
