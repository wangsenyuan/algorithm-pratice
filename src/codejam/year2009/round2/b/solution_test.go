package main

import "testing"

func runSample(t *testing.T, R, C, F int, G []string, expect int) {
	res := solve(R, C, F, G)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", R, C, F, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C, F := 5, 8, 3
	G := []string{
		"........",
		"########",
		"...#.###",
		"####..##",
		"###.##..",
	}
	expect := 2
	runSample(t, R, C, F, G, expect)
}

func TestSample2(t *testing.T) {
	R, C, F := 3, 3, 1
	G := []string{
		"...",
		"###",
		"###",
	}
	expect := 3
	runSample(t, R, C, F, G, expect)
}

func TestSample3(t *testing.T) {
	R, C, F := 2, 2, 1
	G := []string{
		".#",
		"##",
	}
	expect := -1
	runSample(t, R, C, F, G, expect)
}

func TestSample4(t *testing.T) {
	R, C, F := 3, 2, 1
	G := []string{
		"..",
		"#.",
		"..",
	}
	expect := -1
	runSample(t, R, C, F, G, expect)
}
